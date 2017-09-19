package rpcServer

import (
	"xq.goproject.com/commonTool/logTool"
	"fmt"
	"net"
	"sync"
)

//StartServer 开启服务
func StartServer(wg *sync.WaitGroup, serverAddress string) {
	defer func() {
		wg.Done()
	}()

	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		panic(fmt.Errorf("服务器启动失败"))
	}

	//监听服务
	logTool.Log(logTool.Info, fmt.Sprintf("Rpc服务器监听：%s", serverAddress))
	fmt.Println( fmt.Sprintf("Rpc服务器监听：%s", serverAddress))
	go clearClient()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logTool.Log(logTool.Error, fmt.Sprintf("客户端连接错误：%v", err))
			continue
		}

		//处理客户连接
		go HandleConn(conn)
	}
}
