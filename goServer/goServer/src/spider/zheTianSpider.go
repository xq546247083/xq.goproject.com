package spider

import (
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	"xq.goproject.com/goServer/goServer/src/bll/novel"
)

// startZheTianSpider 开始遮天网站抓取
func startZheTianSpider() {
	zhetianSpider := spider.NewSpider(NewMyPageProcesser(), "ZheTianSpider")

	//添加请求地址
	for _, novelConfig := range novel.GetNovelConfigAllList() {
		if novelConfig.SiteName == "遮天小说网" {
			req := request.NewRequest(novelConfig.NovelAddress, "html", novelConfig.NovelName, "GET", "", nil, nil, nil, nil)
			zhetianSpider.AddRequest(req)
		}
	}

	//添加输出
	zhetianSpider.AddPipeline(NewPipelineMysql())

	//设置线程数
	zhetianSpider.SetThreadnum(3)

	zhetianSpider.Run()
}
