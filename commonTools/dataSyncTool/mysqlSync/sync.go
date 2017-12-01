package mysqlSync

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"xq.goproject.com/commonTools/byteTool"
	"xq.goproject.com/commonTools/fileTool"
	"xq.goproject.com/commonTools/logTool"
)

// 同步对象定义
type SyncObject struct {
	// 同步数据的存储路径
	dirPath string

	// 同步数据对象的唯一标识，用于进行重复判断
	identifier string

	// 数据库对象
	dbObj *sql.DB

	// 同步信息对象
	syncInfoObj *syncInfoObject

	// 处理数据写入的文件
	syncDataObj *syncDataObject

	// 同步日志对象
	syncLogObj *syncLogObject
}

// 启动时同步所有数据(然后才能从数据库中查询数据，以免数据丢失)
func (this *SyncObject) syncWhenStart() {
	// 获取文件列表（有序的列表）
	fileList := getDataFileList(this.dirPath)
	syncInfoItemObj := this.syncInfoObj.getSyncInfoItem()

	// 判断是否有文件
	if len(fileList) == 0 {
		return
	}

	// 判断当前文件是否为空，如果为空则将第一个文件赋给它
	if syncInfoItemObj.FilePath == "" {
		syncInfoItemObj.FilePath = fileList[0]
	}
	this.syncInfoObj.setSyncInfoItem(syncInfoItemObj)

	// 开始同步数据
	this.sync()
}

// 循环同步多个文件
func (this *SyncObject) sync() {
	// 获取信息同步项对象
	syncInfoItemObj := this.syncInfoObj.getSyncInfoItem()

	// 开始循环同步
	for {
		// 同步当前文件
		this.syncOneFile(syncInfoItemObj)

		// 当前文件同步完成，记录同步日志
		this.syncLogObj.savaLog(syncInfoItemObj.FilePath)

		// 当前文件同步完成，获取下个文件
		nextFileName := newFileName("", syncInfoItemObj.FilePath)
		filePath := filepath.Join(this.dirPath, nextFileName)
		exist, err := fileTool.IsFileExists(filePath)
		if err != nil {
			panic(err)
		}

		// 如果文件不存在，退出
		if !exist {
			return
		}

		// 重置同步对象
		syncInfoItemObj = newSyncInfoItem(filePath, 0)
		this.syncInfoObj.setSyncInfoItem(syncInfoItemObj)
	}
}

// 同步单个文件
// syncInfoItemObj: 同步项
func (this *SyncObject) syncOneFile(syncInfoItemObj *syncInfoItem) {
	// 打开当前处理文件
	f, err := os.OpenFile(syncInfoItemObj.FilePath, os.O_RDONLY, os.ModePerm|os.ModeTemporary)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncOneFile.os.OpenFile")
		err = fmt.Errorf("%s-Open file:%s failed:%s", prefix, syncInfoItemObj.FilePath, err)
		logTool.LogError(err.Error())
		panic(err)
	}
	defer f.Close()

	// 如果文件是从offset = 0开始同步，需要记录同步信息
	if syncInfoItemObj.Offset == 0 {
		this.syncInfoObj.syncData(syncInfoItemObj)
	}

	for {
		// 1、读取包头内容，并获得包体的长度
		if _, err := f.Seek(syncInfoItemObj.Offset, 0); err != nil {
			prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncOneFile.f.Seek")
			err = fmt.Errorf("%s-Seek offset for header failed:%s", prefix, err)
			logTool.LogError(err.Error())
			panic(err)
		}

		header := make([]byte, 4)
		n, err := f.Read(header)
		if err != nil {
			// 如果读取到文件末尾，判断是否等待
			if err == io.EOF {
				if this.syncDataObj != nil && strings.Contains(syncInfoItemObj.FilePath, this.syncDataObj.bigFileObj.FileName()) {
					time.Sleep(20 * time.Millisecond)
					continue
				}

				// 如果该文件是空文件,同步更新信息
				if syncInfoItemObj.Offset == 0 {
					this.syncInfoObj.syncData(syncInfoItemObj)
				}
				return
			}

			prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncOneFile.f.Read")
			err = fmt.Errorf("%s-Read header failed:%s", prefix, err)
			logTool.LogError(err.Error())
			panic(err)
		}

		if n < con_Header_Length {
			prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncOneFile.f.Read")
			err = fmt.Errorf("%s-There is not enough length of header. Expected:%d, but now got%d bytes.", prefix, con_Header_Length, n)
			logTool.LogError(err.Error())
			panic(err)
		}

		dataLength := byteTool.ByteToInt32(header, byterOrder)

		// 2、读取指定长度的内容
		data := make([]byte, dataLength)
		if _, err := f.Seek(syncInfoItemObj.Offset+con_Header_Length, 0); err != nil {
			prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncOneFile.f.Seek")
			err = fmt.Errorf("%s-Seek offset for data failed:%s", prefix, err)
			logTool.LogError(err.Error())
			panic(err)
		}

		n, err = f.Read(data)
		if err != nil {
			prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncOneFile.f.Read")
			err = fmt.Errorf("%s-Read data failed:%s", prefix, err)
			logTool.LogError(err.Error())
			panic(err)
		}

		if n < int(dataLength) {
			prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncOneFile.f.Read")
			err = fmt.Errorf("%s-There is not enough length of data. Expected:%d, but now got%d bytes.", prefix, con_Header_Length, n)
			logTool.LogError(err.Error())
			panic(err)
		}

		// 3、同步到mysql中
		command := string(data)
		if err := this.syncToMysql(command); err != nil {
			//  发送监控报警
			this.handleSqlError(command)
		}

		// 4、更新syncInfoItem
		syncInfoItemObj.Offset += int64(con_Header_Length + dataLength)
		this.syncInfoObj.syncData(syncInfoItemObj)
	}
}

// 处理sql错误(可能是sql错误，也可能是sql command错误)
// command: sql命令
func (this *SyncObject) handleSqlError(command string) {
	errorFileObj := newErrorFile(this.dirPath, this.identifier)
	defer errorFileObj.delete()

	// 保存当前sql命令
	errorFileObj.saveCommand(command)

	// 循环处理当前命令，直到没有错误
	beginTime := time.Now().Unix()
	var err error
	for {
		// 每隔5分钟，发送警报
		if time.Now().Unix()-beginTime > 5*60 {
			beginTime = time.Now().Unix()
		}

		// 每次循环休眠20秒
		time.Sleep(20 * time.Second)
		command = errorFileObj.readCommand()
		err = this.syncToMysql(command)
		if err != nil {
			continue
		}
		break
	}
}

// 保存数据到本地文件
func (this *SyncObject) save(command string) {
	this.syncDataObj.write(command)
}

// 同步数据到mysql中
func (this *SyncObject) syncToMysql(command string) error {
	_, err := this.dbObj.Exec(command)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "SyncObject.syncToMysql")
		err = fmt.Errorf("%s-%s Update to mysql failed:%s", prefix, command, err)
		logTool.LogError(err.Error())
		fmt.Println("fatal Error:%v", err.Error())
		return err
	}

	return nil
}

// 创新新的mysql同步对象
// dirPath:存放数据的目录
// identifier:当前数据的唯一标识（可以使用数据库表名）
// maxFileSize:每个大文件的最大写入值（单位：Byte）
// survivalTime: 日志存活时间
// dbObj:数据库对象
// 返回值:
// mysql同步对象
func newMysqlSync(dirPath, identifier string, maxFileSize int, survivalTime int, dbObj *sql.DB) *SyncObject {
	dirPath = filepath.Join(dirPath, identifier)

	// 创建更新目录
	err := os.MkdirAll(dirPath, os.ModePerm|os.ModeTemporary)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", identifier, "SyncObject.newMysqlSync.os.MkdirAll")
		err = fmt.Errorf("%s-make dir failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}

	// 构造同步信息对象
	this := &SyncObject{
		dirPath:     dirPath,
		identifier:  identifier,
		dbObj:       dbObj,
		syncInfoObj: newSyncInfoObject(dirPath, identifier),
		syncLogObj:  newSyncLogObject(dirPath, identifier, survivalTime),
	}

	// 先同步完现有数据
	this.syncWhenStart()

	// 构造同步数据对象
	fileName := this.syncInfoObj.getSyncInfoItem().FilePath
	this.syncDataObj = newSyncDataObject(dirPath, identifier, fileName, maxFileSize)

	// 由于同步对象会重新新建一个文件，所以syncInfoItem要重置（避免重复更新）
	syncInfoItemObj := newSyncInfoItem(this.syncDataObj.filefullName(), 0)
	this.syncInfoObj.setSyncInfoItem(syncInfoItemObj)

	// 启动一个新goroutine来负责同步数据
	go func() {
		/* 此处不使用goroutineMgr.Monitor/ReleaseMonitor，因为此处不能捕获panic，需要让外部进程终止执行，
		因为此模块的文件读写为核心逻辑，一旦出现问题必须停止进程，否则会造成脏数据
		*/
		this.sync()
	}()

	return this
}
