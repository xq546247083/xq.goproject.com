package ddzTask

import (
	"time"
)

// 斗地主配置model
type DDZTaskModel struct {
	// 接受邮箱
	ReceiveEmail string

	// 斗地主的账号类型（1：微信，2
	AccountType string

	// 账号
	Account string

	// 当前时间
	CurTime time.Time

	// 奖励条目
	RewardNum int32
}

// 新建斗地主配置model
func NewDDZTaskModel(receiveEmail, accountType, account string, curTime time.Time, rewardNum int32) *DDZTaskModel {
	ddzTaskModel := &DDZTaskModel{
		ReceiveEmail: receiveEmail,
		AccountType:  accountType,
		Account:      account,
		CurTime:      curTime,
		RewardNum:    rewardNum,
	}

	return ddzTaskModel
}
