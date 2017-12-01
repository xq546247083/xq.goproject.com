package main

import (
	"sync"
)

type player struct {
	// 玩家UserID
	UserID string `gorm:"column:UserID;primary_key"`

	// 玩家名称
	UserName string `gorm:"column:UserName"`
}

func (this *player) resetUserName(UserName string) {
	this.UserName = UserName
}

func (this *player) tableUserName() string {
	return "player"
}

func newPlayer(UserID, UserName string) *player {
	return &player{
		UserID:   UserID,
		UserName: UserName,
	}
}

type playerMgr struct {
	playerMap map[string]*player

	mutex sync.Mutex
}

func (this *playerMgr) insert(obj *player) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.playerMap[obj.UserID] = obj
}

func (this *playerMgr) delete(obj *player) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	delete(this.playerMap, obj.UserID)
}

func (this *playerMgr) randomSelect() *player {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for _, obj := range this.playerMap {
		return obj
	}
	return nil
}

func newPlayerMgr() *playerMgr {
	return &playerMgr{
		playerMap: make(map[string]*player),
	}
}
