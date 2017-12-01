package mysqlSync

import (
	"database/sql"
	"fmt"
	"sync"

	"xq.goproject.com/commonTools/logTool"
)

type SyncMgr struct {
	// 同步数据的存储路径
	dirPath string

	// 大文件对象size
	maxFileSize int

	// 数据库对象
	dbObj *sql.DB

	// 同步数据存活时间 (单位：hour)
	survivalTime int

	// 同步对象集合
	syncObjMap map[string]*SyncObject

	// 同步对象锁
	mutex sync.RWMutex
}

// 注册同步对象
// identifier:当前数据的唯一标识（可以使用数据库表名）
func (this *SyncMgr) RegisterSyncObj(identifier string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 判断是否设置了相同的唯一标识，以免弄混淆
	if _, exists := this.syncObjMap[identifier]; exists {
		prefix := fmt.Sprintf("%s-%s", identifier, "SyncMgr.RegisterSyncObj")
		err := fmt.Errorf("%s has already existed, please change another identifier", prefix)
		logTool.LogError(err.Error())
		panic(err)
	}

	syncObj := newMysqlSync(this.dirPath, identifier, this.maxFileSize, this.survivalTime, this.dbObj)
	this.syncObjMap[identifier] = syncObj

	fmt.Println("%s同步对象成功注册进SyncMgr, 当前有%d个同步对象\n", identifier, len(this.syncObjMap))
}

// 保存数据
// identifier:当前数据的唯一标识（可以使用数据库表名）
// command:sql命令
func (this *SyncMgr) Save(identifier string, command string) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	syncObj, exists := this.syncObjMap[identifier]
	if !exists {
		err := fmt.Errorf("syncObj:%s does not existed, please register first", identifier)
		logTool.LogError(err.Error())
		panic(err)
	}

	syncObj.save(command)
}

// 构造同步管理对象
// dirPath: 文件目录
// maxFileSize: 大文件对象大小
// survivalTime: 同步数据存活时间 (单位：hour)
// dbObj: 数据库对象
func NewSyncMgr(dirPath string, maxFileSize int, survivalTime int, dbObj *sql.DB) *SyncMgr {
	return &SyncMgr{
		dirPath:      dirPath,
		maxFileSize:  maxFileSize,
		survivalTime: survivalTime,
		dbObj:        dbObj,
		syncObjMap:   make(map[string]*SyncObject),
	}
}
