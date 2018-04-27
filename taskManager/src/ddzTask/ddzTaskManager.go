package	ddzTask

import (
	"fmt"
	"time"
	"strings"

	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/emailTool"
)

var (
	splitChar ="|"
)

// 开始计算今天的选择
func CaclReward(){
	ddzTaskList:=getDDZTask()
	_=ddzTaskList

	// 循环计算奖励,构建邮件字符串
	dtNow:=time.Now()

	// 按发送qq分组
	sendQQList:=make([]int64,0,2)
	for _,ddzTask:=range ddzTaskList{
		// 判断是否存在，如果不存在才添加
		isExist:=false
		for _,qqItem:=range sendQQList{
			if qqItem==ddzTask.SendQQ{
				isExist=true
			}
		}

		if!isExist{
			sendQQList=append(sendQQList,ddzTask.SendQQ)
		}
	}

	// 给发送qq组装对应的斗地主信息
	for _,qqItem:=range sendQQList{
		ddzEmailContentStr:="美女你好，你今天的运势是:<br/>"
		for _,ddzTask:=range ddzTaskList{
			if ddzTask.SendQQ==qqItem{
				daySpan:=int32(dtNow.Sub(ddzTask.CurTime).Hours()/24)
				curRewardNum:= (ddzTask.RewardNum+ daySpan)%3+1

				ddzEmailContentStr+= fmt.Sprintf("QQ:%d&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;",ddzTask.QQ)
				ddzEmailContentStr+=fmt.Sprintf("请选择:%d<br/>",curRewardNum)
			}
		}

		// 构建收件人
		mailTo:=make([]string,0,1)
		mailTo=append(mailTo, fmt.Sprintf("%d@qq.com",qqItem))
		emailTool.SendMail(mailTo,"斗地主专用",ddzEmailContentStr,true,nil)
	}
}

// 获取斗地主任务列表
func getDDZTask()  []*DDZTaskModel{
	// 结果
	result:=make([]*DDZTaskModel,0,2)

	// 获取任务配置字符串
	taskStr:= configTool.DDZTask

	// 拆分，并解析对应的任务
	taskStrList:= strings.Split(taskStr,splitChar)
	for _,taskItemStr:=range taskStrList{
		itemList:= stringTool.SplitToInt64List(taskItemStr)
		if len(itemList)!=4{
			panic("斗地主任务字符串配置错误。案例：546247083,100000,1|295787943,100000,1")
		}

		curTime:=time.Unix(itemList[2],0)
		ddzTaskModel:=NewDDZTaskModel(itemList[0],itemList[1],curTime ,int32(itemList[3]))
		result=append(result,ddzTaskModel)
	}

	return result
}