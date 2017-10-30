package spider

import (
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
)

// Start 开始抓取任务
func Start() {
	mainSpider := spider.NewSpider(NewMyPageProcesser(), "TaskName")

	//添加请求地址
	req := request.NewRequest("http://m.zhetian.org/1361/list/", "html", "圣墟", "GET", "", nil, nil, nil, nil)
	mainSpider.AddRequest(req)

	//添加输出
	mainSpider.AddPipeline(NewPipelineMysql())

	//设置线程数
	mainSpider.SetThreadnum(3)

	mainSpider.Run()
}
