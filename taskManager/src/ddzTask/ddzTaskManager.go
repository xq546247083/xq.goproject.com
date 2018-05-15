package ddzTask

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/emailTool"
)

var (
	splitChar  = "|"
	splitChar2 = ","
)

// 开始计算今天的选择
func CaclReward() {
	ddzTaskList := getDDZTask()
	_ = ddzTaskList

	// 循环计算奖励,构建邮件字符串
	dtNow := time.Now()

	// 按接受邮箱分组
	receiveList := make([]string, 0, 2)
	for _, ddzTask := range ddzTaskList {
		// 判断是否存在，如果不存在才添加
		isExist := false
		for _, reveiveEmail := range receiveList {
			if reveiveEmail == ddzTask.ReceiveEmail {
				isExist = true
			}
		}

		if !isExist {
			receiveList = append(receiveList, ddzTask.ReceiveEmail)
		}
	}

	// 给发送qq组装对应的斗地主信息
	for _, reveiveEmail := range receiveList {
		ddzEmailContentStr := "亲爱的娟，你今天的运势是:<br/>"
		for _, ddzTask := range ddzTaskList {
			if ddzTask.ReceiveEmail == reveiveEmail {
				daySpan := int32(dtNow.Sub(ddzTask.CurTime).Hours() / 24)
				curRewardNum := (ddzTask.RewardNum+daySpan)%3 + 1

				ddzEmailContentStr += fmt.Sprintf("类型:%s&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;", ddzTask.AccountType)
				ddzEmailContentStr += fmt.Sprintf("账号:%s&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;", ddzTask.Account)
				ddzEmailContentStr += fmt.Sprintf("请选择:%d<br/>", curRewardNum)
			}
		}

		// 构建收件人
		mailTo := make([]string, 0, 1)
		mailTo = append(mailTo, reveiveEmail)
		emailTool.SendMail(mailTo, "斗地主啊斗地主", ddzEmailContentStr, true, nil)
	}
}

// 获取斗地主任务列表
func getDDZTask() []*DDZTaskModel {
	// 结果
	result := make([]*DDZTaskModel, 0, 2)

	// 获取任务配置字符串
	taskStr := configTool.DDZTask

	// 拆分，并解析对应的任务
	taskStrList := strings.Split(taskStr, splitChar)
	for _, taskItemStr := range taskStrList {
		itemList := strings.Split(taskItemStr, splitChar2)
		if len(itemList) != 5 {
			panic("斗地主任务字符串配置错误。案例：2545625776@qq.com,微信,546247083,1514736000,1|2545625776@qq.com,QQ,546247083,1514736000,2")
		}

		// 获取配置
		curTimeUnix, err := strconv.ParseInt(itemList[3], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("斗地主任务字符串配置错误。CurTime:【%s】应该为数字！", itemList[3]))
		}
		curTime := time.Unix(curTimeUnix, 0)

		rewardNum, err := strconv.Atoi(itemList[4])
		if err != nil {
			panic(fmt.Sprintf("斗地主任务字符串配置错误。RewardNum：【%s】应该为数字！", itemList[4]))
		}

		ddzTaskModel := NewDDZTaskModel(itemList[0], itemList[1], itemList[2], curTime, int32(rewardNum))
		result = append(result, ddzTaskModel)
	}

	return result
}
