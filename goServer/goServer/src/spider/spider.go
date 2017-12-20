package spider

import "time"

// Start 开始抓取数据
func Start() {
	go startWowShopSpider()

	//  每分钟执行一次
	for {
		// 获取当前时间的日期
		now := time.Now()
		if (now.Hour() == 5 || now.Hour() == 11 || now.Hour() == 17 || now.Hour() == 22) && now.Minute() == 0 {
			// go startZheTianSpider()
		}

		remainTime := time.Duration(60 - time.Now().Second())
		time.Sleep(remainTime * time.Second)
	}
}
