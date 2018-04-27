package	ddzTask

import (
	"time"
)

// 斗地主任务model
type DDZTaskModel struct{
	// SendQQ
	SendQQ int64

	// qq
	QQ int64

	// 当前时间
	CurTime time.Time

	// 奖励条目
	RewardNum int32
}

// 新建斗地主model
func NewDDZTaskModel(sendQQ int64,qq int64,curTime time.Time,rewardNum int32)*DDZTaskModel{
	ddzTaskModel:=&DDZTaskModel{
		SendQQ:sendQQ,
		QQ:qq,
		CurTime:curTime,
		RewardNum:rewardNum,
	}
		
	return ddzTaskModel
}