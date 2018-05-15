package task

import (
	"time"

	"xq.goproject.com/taskManager/src/ddzTask"
)

// 开启任务监控
func Moniter() {
	//  每分钟执行一次
	for {
		// 获取当前时间的日期
		now := time.Now()
		if now.Hour() == 1 && now.Minute() == 0 {
			ddzTask.CaclReward()
		}

		remainTime := time.Duration(60 - time.Now().Second())
		time.Sleep(remainTime * time.Second)
	}
}
