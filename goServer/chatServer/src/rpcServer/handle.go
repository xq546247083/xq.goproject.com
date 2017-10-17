package rpcServer

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"xq.goproject.com/commonTools/goroutineTool"
	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/goServer/goServerModel/src/common"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
)

//HandleConn 处理客户端连接
func HandleConn(conn net.Conn) {
	clientObj := newClient(conn)
	registerClient(clientObj)

	goroutineTool.Operate("HandleConn", common.AddOperate)
	defer goroutineTool.Operate("HandleConn", common.ReduceOperate)
	//是否发送数据的通道
	ch := make(chan bool, 1)
	go handleSendData(clientObj, ch)

	//释放链接
	defer func() {
		clientObj.Quit()
		unRegisterClient(clientObj.id)

		//关闭发送数据
		ch <- true
	}()

	//监听客户端程序
	for {
		readData := make([]byte, 1024)

		//阻塞读取数据
		len, err := conn.Read(readData)
		if err != nil {
			//logTool.Log(logTool.Error, fmt.Sprintf("接受消息出错,客户端为：%v\n，错误：%s\n，消息长度：%d\n", clientObj, err, len))
			break
		}

		clientObj.appendReceiveData(readData[:len])
		handleReceiveData(clientObj)
	}
}

//handleReceiveData 处理客户端数据
func handleReceiveData(clientObj *Client) {
	for {
		// 获取有效的消息
		message, exists := clientObj.getReceiveData()
		if !exists {
			break
		}

		//处理数据，如果长度为0则表示心跳包；否则处理请求内容
		if len(message) == 0 {
			continue
		} else {
			handRequest(clientObj, message)
		}
	}
}

//handRequest 处理请求
func handRequest(clientObj *Client, message []byte) {
	responseObj := rpcServerObject.NewResponseObject()
	var requestObj rpcServerObject.RequestObject

	// 解析请求字符串
	if err := json.Unmarshal(message, &requestObj); err != nil {
		logTool.Log(logTool.Error, fmt.Sprintf("反序列化出错，错误信息为：%s", err))

		ResponseResult(clientObj, responseObj.SetResultStatus(rpcServerObject.DataError), ConHighPriority)
		return
	}

	logTool.Log(logTool.Debug, "RPC服务器接受到请求："+string(message))

	//先判断用户请求,如果通过，再调用方法
	requestObj.RequestInfo["Client"] = clientObj
	callFunction(&requestObj)
}

//handSendData 处理发送数据
func handleSendData(clientObj *Client, ch chan bool) {
	goroutineTool.Operate("handleSendData", common.AddOperate)
	defer goroutineTool.Operate("handleSendData", common.ReduceOperate)

	for {
		select {
		case <-ch:
			//收到消息，客户端退出，退出线程
			return
		default:
			//没有退出，默认一直循环处理消息
			for {

				if clientObj.connStatus == connClose {
					break
				}

				//处理消息
				if responese, exist := clientObj.getSendData(); exist {
					if err := clientObj.sendMessage(responese); err != nil {
						//如果发送消息失败，退出发送线程
						return
					}
				} else {
					//如果没有数据可发送，退出循环
					break
				}
			}

			time.Sleep(50 * time.Millisecond)
		}
	}
}
