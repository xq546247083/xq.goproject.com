package spider

import (
	"fmt"
	"strings"

	"xq.goproject.com/commonTools/logTool"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
)

// MyPageProcesser 页面处理结构
type MyPageProcesser struct {
}

// NewMyPageProcesser 新建页面处理
func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

// Process Process 处理爬到的页面
func (thisObj *MyPageProcesser) Process(p *page.Page) {
	query := p.GetHtmlParser()
	var urls []string
	query.Find("h3[class='repo-list-name'] a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		urls = append(urls, "http://github.com/"+href)
	})
	// these urls will be saved and crawed by other coroutines.
	p.AddTargetRequests(urls, "html")

	name := query.Find(".entry-title .author").Text()
	name = strings.Trim(name, " \t\n")
	repository := query.Find(".entry-title .js-current-repository").Text()
	repository = strings.Trim(repository, " \t\n")
	//readme, _ := query.Find("#readme").Html()
	if name == "" {
		p.SetSkip(true)
	}
	// the entity we want to save by Pipeline
	p.AddField("author", name)
	p.AddField("project", repository)
	//p.AddField("readme", readme)

	logTool.LogDebug(p.GetBodyStr())
}

// Finish 完成爬虫任务
func (thisObj *MyPageProcesser) Finish() {
	fmt.Println("任务完成")
	logTool.LogDebug("任务完成")
}
