package player

import (
	"time"
)

//Player 定义玩家对象
type Player struct {
	// 玩家ID
	ID string

	// 玩家名称
	Name string

	// 玩家等级
	Lv int

	// 玩家Vip等级
	Vip int

	// 合作商ID
	PartnerID int

	// 服务器ID
	ServerID int

	// 玩家公会ID
	UnionID string

	// 额外透传信息
	ExtraMsg string

	// 服务器名称
	ServerName string

	// 服务器组ID
	ServerGroupID int

	//注册时间
	RegisterTime time.Time `json:"-"`

	//登录时间
	LoginTime time.Time `json:"-"`

	//是否封号
	IsForbidden bool `json:"-"`

	//禁言结束时间
	SilentEndTime time.Time `json:"-"`

	// 上次发送消息的时间懂戳
	LastSendTime int64 `json:"-"`

	// 消息历史
	HistoryMessageList []map[rune]bool `json:"-"`

	// 客户端
	ClientID int32
}

//InitPlayer 初始化一个玩家对象
func InitPlayer(ID, name string, lv, vip, partnerID, serverID int, unionID string, extraMsg string) *Player {
	return &Player{
		ID:                 ID,
		Name:               name,
		Lv:                 lv,
		Vip:                vip,
		PartnerID:          partnerID,
		ServerID:           serverID,
		UnionID:            unionID,
		ExtraMsg:           extraMsg,
		RegisterTime:       time.Now(),
		LoginTime:          time.Now(),
		IsForbidden:        false,
		SilentEndTime:      time.Now(),
		LastSendTime:       0,
		HistoryMessageList: make([]map[rune]bool, 0, 16),
	}
}

//NewPlayer 使用现有数据构造一个新的玩家对象
func NewPlayer(id, name string, lv, vip, partnerID, serverID int, unionID, extraMsg string, registerTime, loginTime time.Time, isForbidden bool, silentEndTime time.Time) *Player {
	return &Player{
		ID:                 id,
		Name:               name,
		Lv:                 lv,
		Vip:                vip,
		PartnerID:          partnerID,
		ServerID:           serverID,
		UnionID:            unionID,
		ExtraMsg:           extraMsg,
		RegisterTime:       registerTime,
		LoginTime:          loginTime,
		IsForbidden:        isForbidden,
		SilentEndTime:      silentEndTime,
		LastSendTime:       0,
		HistoryMessageList: make([]map[rune]bool, 0, 16),
	}
}

//NewEmptyPlayer 使用现有数据构造一个新的玩家对象
func NewEmptyPlayer(id string, clientID int32) *Player {
	return &Player{
		ID:       id,
		ClientID: clientID,
	}
}
