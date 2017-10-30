package spider

import (
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
)

// Start 开始爬虫任务
func Start() {
	mainSpider := spider.NewSpider(NewMyPageProcesser(), "TaskName")

	mainSpider.AddUrl("https://github.com/hu17889?tab=repositories", "html")
	mainSpider.AddPipeline(pipeline.NewPipelineConsole())
	mainSpider.SetThreadnum(3)

	mainSpider.Run()
}
