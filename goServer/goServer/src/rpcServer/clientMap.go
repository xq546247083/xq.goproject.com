package rpcServer

import (
	"sync"
)

var (
	clientMap = make(map[int32]*Client, 1024)

	clientMutex sync.RWMutex
)

//registerClient 注册客户端
func registerClient(clientObj *Client) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	clientMap[clientObj.id] = clientObj
}

//unRegisterClient 注销客户端
func unRegisterClient(clientID int32) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	delete(clientMap, clientID)
}

//GetClient 获取客户端
func GetClient(clientID int32) *Client {
	clientMutex.RLock()
	defer clientMutex.RUnlock()

	if clientObj, ok := clientMap[clientID]; ok {
		return clientObj
	}

	return nil
}

//getClientCount 获取客户端的数量
// 返回值：客户端数量
func getClientCount() int {
	clientMutex.RLock()
	defer clientMutex.RUnlock()

	return len(clientMap)
}

//getObsoleteClient 获取过时的客户端
func getObsoleteClient() (result []*Client) {
	clientMutex.RLock()
	defer clientMutex.RUnlock()
	for _, item := range clientMap {
		if item.isObsolete() {
			result = append(result, item)
		}
	}

	return
}
