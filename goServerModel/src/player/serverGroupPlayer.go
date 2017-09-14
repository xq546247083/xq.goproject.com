package player

import (
	"sync"
)

//ServerGroupPlayer 服务器组玩家对象
type ServerGroupPlayer struct {
	// 服务器组ID
	serverGroupID int

	// 玩家集合
	//key:玩家id
	//value：玩家对象
	playerMap map[string]*Player

	// 锁对象
	mutex sync.RWMutex
}

//GetPlayerList 获取玩家列表
func (serverGroupPlayerObj *ServerGroupPlayer) GetPlayerList() (playerList []*Player) {
	serverGroupPlayerObj.mutex.RLock()
	defer serverGroupPlayerObj.mutex.RUnlock()

	for _, playerObj := range serverGroupPlayerObj.playerMap {
		playerList = append(playerList, playerObj)
	}

	return
}

//GetPlayerListInUnion 获取玩家列表InUnion
func (serverGroupPlayerObj *ServerGroupPlayer) GetPlayerListInUnion(unionID string) (playerList []*Player) {
	serverGroupPlayerObj.mutex.RLock()
	defer serverGroupPlayerObj.mutex.RUnlock()

	for _, playerObj := range serverGroupPlayerObj.playerMap {
		if playerObj.UnionID == unionID {
			playerList = append(playerList, playerObj)
		}
	}

	return
}

//AddPlayer 添加玩家
func (serverGroupPlayerObj *ServerGroupPlayer) AddPlayer(playerObj *Player) {
	serverGroupPlayerObj.mutex.Lock()
	defer serverGroupPlayerObj.mutex.Unlock()

	serverGroupPlayerObj.playerMap[playerObj.ID] = playerObj
}

//DeletePlayer 删除玩家
func (serverGroupPlayerObj *ServerGroupPlayer) DeletePlayer(playerObj *Player) {
	serverGroupPlayerObj.mutex.Lock()
	defer serverGroupPlayerObj.mutex.Unlock()

	delete(serverGroupPlayerObj.playerMap, playerObj.ID)
}

//NewServerGroupPlayer 新建一个服务器玩家组对象
func NewServerGroupPlayer(_serverGroupID int) *ServerGroupPlayer {
	return &ServerGroupPlayer{
		serverGroupID: _serverGroupID,
		playerMap:     make(map[string]*Player, 1024),
	}
}
