package rpcServer

import (
	"xq.goproject.com/goServerModel/src/rpcServerObject"
	"xq.goproject.com/commonTool/byteTool"
	"xq.goproject.com/commonTool/intTool"
	"xq.goproject.com/commonTool/logTool"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// 包头的长度
	conHeaderLength = 4

	//id的长度
	conIDLength
)

var (
	//全局自增id
	globalID int32 = 1

	//字节顺序
	byteOrder = binary.LittleEndian
)

//Client 客户端
type Client struct {
	//id
	id int32

	//激活时间
	activeTime time.Time

	//连接
	conn net.Conn

	//连接状态
	connStatus connStatus

	//当前客户端的锁
	mutex sync.Mutex

	//接受的数据
	receiveData []byte

	//待发送的数据
	sendData []*rpcServerObject.ResponseObject

	//待发送的慢数据
	sendSlowData []*rpcServerObject.ResponseObject
}

//新建客户端
func newClient(_conn net.Conn) *Client {
	getIncrementID := func() int32 {
		atomic.AddInt32(&globalID, 1)
		return globalID
	}

	return &Client{
		id:           getIncrementID(),
		activeTime:   time.Now(),
		conn:         _conn,
		connStatus:   connOpen,
		receiveData:  make([]byte, 0, 1024),
		sendData:     make([]*rpcServerObject.ResponseObject, 0, 16),
		sendSlowData: make([]*rpcServerObject.ResponseObject, 0, 16),
	}
}

//Quit 客户端退出
func (clientObj *Client) Quit() {
	clientObj.conn.Close()
	clientObj.connStatus = connClose
}

//获取远程地址（IP_Port）
func (clientObj *Client) getRemoteAddr() string {
	items := strings.Split(clientObj.conn.RemoteAddr().String(), ":")

	return fmt.Sprintf("%s_%s", items[0], items[1])
}

//获取远程地址（IP）
func (clientObj *Client) getRemoteShortAddr() string {
	items := strings.Split(clientObj.conn.RemoteAddr().String(), ":")

	return items[0]
}

//GetID 获取id
func (clientObj *Client) GetID() int32 {
	return clientObj.id
}

//appendReceiveData 追加数据
func (clientObj *Client) appendReceiveData(data []byte) {
	clientObj.receiveData = append(clientObj.receiveData, data...)
	clientObj.activeTime = time.Now()
}

//getReceiveData 获取接受数据,并处理掉获取的消息
func (clientObj *Client) getReceiveData() ([]byte, bool) {
	var data = clientObj.receiveData
	if len(data) < conHeaderLength {
		return nil, false
	}

	contentLength := byteTool.ByteToInt32(data[:conHeaderLength], byteOrder)
	if contentLength+conHeaderLength > int32(len(data)) {
		return nil, false
	}

	//获取内容
	content := data[conHeaderLength : contentLength+conHeaderLength]
	clientObj.receiveData = data[contentLength+conHeaderLength:]

	return content, true
}

//appendSendData 追加发送的数据
// sendDataItemObj:待发送数据项
// priority:优先级
// 返回值：无
func (clientObj *Client) appendSendData(responseObj *rpcServerObject.ResponseObject, priority Priority) {
	if priority == ConLowPriority {
		clientObj.sendSlowData = append(clientObj.sendSlowData, responseObj)
	} else {
		clientObj.sendData = append(clientObj.sendData, responseObj)
	}
}

//getSendData 获取发送数据,并处理掉发送的消息
func (clientObj *Client) getSendData() (responseObject *rpcServerObject.ResponseObject, flag bool) {
	clientObj.mutex.Lock()
	defer clientObj.mutex.Unlock()

	if len(clientObj.sendData) > 0 {
		responseObject = clientObj.sendData[0]
		clientObj.sendData = clientObj.sendData[1:]

		return responseObject, true
	}

	return nil, false
}

//getSendLowData 获取发送低权限数据，并处理掉发送的消息
func (clientObj *Client) getSendSlowData() (*rpcServerObject.ResponseObject, bool) {
	clientObj.mutex.Lock()
	defer clientObj.mutex.Unlock()

	if len(clientObj.sendSlowData) > 0 {
		clientObj.sendSlowData = clientObj.sendSlowData[1:]

		return clientObj.sendSlowData[0], true
	}

	return nil, false
}

//sendMessage 发送消息
func (clientObj *Client) sendMessage(responseObject *rpcServerObject.ResponseObject) (err error) {
	//序列化发送的数据
	content, err := json.Marshal(responseObject)
	if err != nil {
		logTool.Log(logTool.Error, fmt.Sprintf("序列化要发送的消息错误，message:%s,err:%s", content, err))
		return
	}

	//把文字长度写进头
	contentLength := len(content)
	head := intTool.IntToByte(contentLength, byteOrder)

	message := append(head, content...)
	if _, err = clientObj.conn.Write(message); err != nil {
		logTool.Log(logTool.Error, fmt.Sprintf("发送消息错误，message:%s,err:%s", message, err))
		return
	}else{
		logTool.Log(logTool.Debug, "RPC服务器发送数据："+string(message))
	}

	return
}

//isObsolete 判断客户端是否超时（超过300秒不活跃算作超时）
// 返回值：是否超时
func (clientObj *Client) isObsolete() bool {
	return time.Now().Unix() > clientObj.activeTime.Add(300*time.Second).Unix()
}

//ResponseResult 响应结果
func ResponseResult(clientObj *Client, responseObject *rpcServerObject.ResponseObject, priority Priority) {
	if clientObj != nil {
		clientObj.mutex.Lock()
		defer clientObj.mutex.Unlock()

		clientObj.appendSendData(responseObject, priority)
	}
}
