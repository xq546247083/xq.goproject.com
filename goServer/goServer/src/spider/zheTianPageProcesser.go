package spider

import (
	"fmt"
	"strings"
	"time"

	"xq.goproject.com/goServer/goServer/src/bll/novel"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"xq.goproject.com/commonTools/stringTool"
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
	novelName := p.GetUrlTag()
	query := p.GetHtmlParser()
	var reqs []*request.Request

	//如果是下一页，继续爬
	query.Find("a[class='ptm-btn ptm-btn-primary ptm-btn-block ptm-btn-outlined']").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "下一页" {
			href, _ := s.Attr("href")
			req := request.NewRequest(stringTool.GetURLDomainName(p.GetRequest().GetUrl())+href, "html", p.GetUrlTag(), "GET", "", nil, nil, nil, nil)
			reqs = append(reqs, req)
		}
	})

	//如果是章节页面，继续爬
	query.Find("li[class='ptm-list-view-cell'] a").Each(func(i int, s *goquery.Selection) {
		//如果数据库不存在该章节
		if !novel.IsExisItems(novelName, s.Text()) {
			href, _ := s.Attr("href")
			req := request.NewRequest(stringTool.GetURLDomainName(p.GetRequest().GetUrl())+href, "html", p.GetUrlTag(), "GET", "", nil, nil, nil, nil)
			reqs = append(reqs, req)
		}
	})

	//如果是换源页面，继续爬
	query.Find("div[class='pt-name'] a[class='ptm-text-grey']").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		req := request.NewRequest(stringTool.GetURLDomainName(p.GetRequest().GetUrl())+href, "html", p.GetUrlTag(), "GET", "", nil, nil, nil, nil)
		reqs = append(reqs, req)
	})

	p.AddTargetRequestsWithParams(reqs)

	//处理页面数据
	title := strings.Trim(query.Find("h1[class='title']").Text(), " \t\n")
	source := strings.Trim(query.Find("div[class='d_out'] div[class='d_menu']").Text(), " \t\n")
	htmlStr, errHTML := query.Find("div[class='articlecon']").Html()
	if htmlStr == "" || title == "" || errHTML != nil {
		p.SetSkip(true)
	} else {
		p.AddField("name", novelName)
		p.AddField("title", title)
		p.AddField("source", source)
		p.AddField("content", htmlStr)
	}
}

// Finish 完成爬虫任务
func (thisObj *MyPageProcesser) Finish() {
	fmt.Println("遮天网站抓取完成")

	//6个小时候，继续抓取
	time.Sleep(time.Hour * 6)
	go startZheTianSpider()
}
