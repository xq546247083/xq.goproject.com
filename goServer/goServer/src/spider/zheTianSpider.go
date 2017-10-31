package spider

import (
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
)

// startZheTianSpider 开始遮天网站抓取
func startZheTianSpider() {
	zhetianSpider := spider.NewSpider(NewMyPageProcesser(), "ZheTianSpider")

	//添加请求地址
	req := request.NewRequest("http://m.zhetian.org/1361/list/", "html", "圣墟", "GET", "", nil, nil, nil, nil)
	zhetianSpider.AddRequest(req)

	//添加输出
	zhetianSpider.AddPipeline(NewPipelineMysql())

	//设置线程数
	zhetianSpider.SetThreadnum(3)

	zhetianSpider.Run()
}
