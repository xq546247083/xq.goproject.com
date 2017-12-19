package spider

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/goServer/src/bll/novel"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

//------------------------------------------------------抓取开始代码-----------------------------------------------
// startZheTianSpider 开始遮天网站抓取
func startZheTianSpider() {
	zhetianSpider := spider.NewSpider(NewZheTianProcesser(), "ZheTianSpider")

	//添加请求地址
	for _, novelConfig := range novel.GetNovelConfigAllList() {
		if novelConfig.SiteName == "遮天小说网" {
			req := request.NewRequest(novelConfig.NovelAddress, "html", novelConfig.NovelName, "GET", "", nil, nil, nil, nil)
			zhetianSpider.AddRequest(req)
		}
	}

	//添加输出
	zhetianSpider.AddPipeline(NewZheTianMysql())

	//设置线程数
	zhetianSpider.SetThreadnum(1)

	zhetianSpider.Run()
}

//------------------------------------------------------页面处理-----------------------------------------------

// ZheTianProcesser 页面处理结构
type ZheTianProcesser struct {
}

// NewZheTianProcesser 新建页面处理
func NewZheTianProcesser() *ZheTianProcesser {
	return &ZheTianProcesser{}
}

// Process Process 处理爬到的页面
func (thisObj *ZheTianProcesser) Process(p *page.Page) {
	urlTag := p.GetUrlTag()
	query := p.GetHtmlParser()
	var reqs []*request.Request

	//如果是下一页，继续爬
	query.Find("a[class='ptm-btn ptm-btn-primary ptm-btn-block ptm-btn-outlined']").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "下一页" {
			href, _ := s.Attr("href")
			if strings.Contains(href, ";") {
				return
			}

			req := request.NewRequest(stringTool.GetURLDomainName(p.GetRequest().GetUrl())+href, "html", p.GetUrlTag(), "GET", "", nil, nil, nil, nil)
			reqs = append(reqs, req)
		}
	})

	//如果是章节页面，继续爬
	query.Find("li[class='ptm-list-view-cell'] a").Each(func(i int, s *goquery.Selection) {
		//如果数据库不存在该章节
		if !novel.IsExisItems(urlTag, s.Text()) {
			href, _ := s.Attr("href")
			if strings.Contains(href, ";") {
				return
			}

			req := request.NewRequest(stringTool.GetURLDomainName(p.GetRequest().GetUrl())+href, "html", fmt.Sprintf("%s,%s", p.GetUrlTag(), s.Text()), "GET", "", nil, nil, nil, nil)
			reqs = append(reqs, req)
		}
	})

	//如果是换源页面，继续爬
	query.Find("span[class='pt-name'] a[class='ptm-text-grey']").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.Contains(href, ";") {
			return
		}

		req := request.NewRequest(stringTool.GetURLDomainName(p.GetRequest().GetUrl())+href, "html", p.GetUrlTag(), "GET", "", nil, nil, nil, nil)
		reqs = append(reqs, req)
	})

	p.AddTargetRequestsWithParams(reqs)

	//处理页面数据
	urlTags := strings.Split(urlTag, ",")
	if len(urlTags) != 2 {
		p.SetSkip(true)
		return
	}

	//title := strings.Trim(query.Find("h1[class='title']").Text(), " \t\n")
	source := strings.Trim(query.Find("div[class='d_out'] div[class='d_menu']").Text(), " \t\n")
	htmlStr, errHTML := query.Find("div[class='articlecon']").Html()
	if urlTags[0] == "" || urlTags[1] == "" || errHTML != nil || htmlStr == "<p></p>" || htmlStr == "" {
		p.SetSkip(true)
	} else {
		p.AddField("name", urlTags[0])
		p.AddField("title", urlTags[1])
		p.AddField("source", source)
		p.AddField("content", htmlStr)
	}
}

// Finish 完成爬虫任务
func (thisObj *ZheTianProcesser) Finish() {
	fmt.Println("遮天网站抓取完成")
}

//------------------------------------------------------数据持久处理-----------------------------------------------

// ZheTianMysql mysql输出持久层
type ZheTianMysql struct {
}

// NewZheTianMysql 新建mysql输出持久层
func NewZheTianMysql() *ZheTianMysql {
	return &ZheTianMysql{}
}

// Process 处理获得的数据
func (thisObj *ZheTianMysql) Process(items *page_items.PageItems, t com_interfaces.Task) {
	name := items.GetAll()["name"]
	title := items.GetAll()["title"]
	source := strings.Replace(strings.Replace(items.GetAll()["source"], "↓", "", -1), " ", "", -1)
	content := items.GetAll()["content"]

	novel := model.NewNovel(name, title, source, content)
	dal.NovelDALObj.SaveInfo(novel, nil)
}
