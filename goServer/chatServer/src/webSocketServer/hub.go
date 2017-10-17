// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webSocketServer

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

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		}
	}
}

// SendMessage 给客户端发送消息
func SendMessage(userName string, message []byte) {
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

// BroadMessage 广播消息
func BroadMessage(message []byte) {
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
