// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webSocketServer

import (
	"encoding/json"

	"xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func runHub() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}

			broadClients()
		}
	}
}

// SetOnlineStatus 设置在线状态
func SetOnlineStatus(userName string, flag bool) {
	for client := range hub.clients {
		if userName == client.userName {
			hub.clients[client] = flag
		}
	}
}

// IsOnline 玩家是否在线
func IsOnline(userName string) bool {
	for client, flag := range hub.clients {
		if userName == client.userName && flag {
			return true
		}
	}

	return false
}

//  broadClients 广播客户端
func broadClients() {
	requestObj := webSocketServerObject.NewRequestObject()
	requestObj.MethodName = "BroadClients"
	callFunction(requestObj)
}

// BroadMessage 广播消息
func BroadMessage(responseObject *webSocketServerObject.ResponseObject) {
	message, err := json.Marshal(responseObject)
	if err != nil {
		//返回对象反序列化失败，只能返回空数据
		message = []byte("")
	}

	for c, flag := range hub.clients {
		if flag {
			select {
			case c.send <- message:
			default:
				close(c.send)
				c.hub.unregister <- c
			}
		}
	}
}

// SendMessage 给客户端发送消息
func SendMessage(userName string, responseObject *webSocketServerObject.ResponseObject) {
	message, err := json.Marshal(responseObject)
	if err != nil {
		//返回对象反序列化失败，只能返回空数据
		message = []byte("")
	}

	for c, flag := range hub.clients {
		if flag && c.userName == userName {
			select {
			case c.send <- message:
			default:
				close(c.send)
				c.hub.unregister <- c
			}
		}
	}
}
