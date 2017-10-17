// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webSocketServer

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"xq.goproject.com/commonTools/logTool"
)

var (
	//hub
	hub = newHub()
)

// StartServer 开启服务
func StartServer(wg *sync.WaitGroup, serverAddress string) {
	defer func() {
		wg.Done()
	}()

	//地址
	var addr = flag.String("addr", serverAddress, "http service address")
	flag.Parse()

	//监听服务
	logTool.Log(logTool.Info, fmt.Sprintf("webSocketServer服务器监听：%s", serverAddress))
	fmt.Println(fmt.Sprintf("webSocketServer服务器监听：%s", serverAddress))

	go hub.run()

	// 处理事件
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	//监听地址
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
