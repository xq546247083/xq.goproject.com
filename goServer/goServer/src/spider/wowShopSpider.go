package spider

import (
	"fmt"
	"net/http"

	"xq.goproject.com/commonTools/logTool"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	"xq.goproject.com/commonTools/emailTool"
)

//------------------------------------------------------抓取开始代码-----------------------------------------------
// startWowShopSpider 开始魔兽商店数据抓取
func startWowShopSpider() {
	wowShopSpider := spider.NewSpider(NewWowShopProcesser(), "WowShopSpider")

	//添加请求地址
	header := make(http.Header)
	// header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER 1.1")

	req := request.NewRequest("https://shop.battlenet.com.cn/zh-cn/family/world-of-warcraft", "html", "WowShopSpider", "GET", "", header, nil, nil, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER 1.1")
	wowShopSpider.AddRequest(req)

	//添加输出
	wowShopSpider.AddPipeline(NewWowShopEmail())

	//设置线程数
	wowShopSpider.SetThreadnum(1)

	wowShopSpider.Run()
}

//------------------------------------------------------页面处理-----------------------------------------------

// WowShopProcesser 页面处理结构
type WowShopProcesser struct {
}

// NewWowShopProcesser 新建页面处理
func NewWowShopProcesser() *WowShopProcesser {
	return &WowShopProcesser{}
}

// Process Process 处理爬到的页面
func (thisObj *WowShopProcesser) Process(p *page.Page) {
	query := p.GetHtmlParser()

	logTool.LogError(p.GetBodyStr())
	// 阵营转换服务
	query.Find("span[class='full']").Each(func(i int, s *goquery.Selection) {
		sss := s.Text()
		_ = sss
	})

	price, errHTML := query.Find("a[href='/zh-cn/product/world-of-warcraft-service-faction-change']").ChildrenFiltered("span[class='price']").ChildrenFiltered("span").Html()
	_ = errHTML
	p.AddField("Name", "阵营转换服务")
	p.AddField("Price", price)
}

// Finish 完成爬虫任务
func (thisObj *WowShopProcesser) Finish() {
	fmt.Println("魔兽商店数据拉取完成")
}

//------------------------------------------------------数据持久处理-----------------------------------------------

// WowShopEmail 邮件输出
type WowShopEmail struct {
}

// NewWowShopEmail 邮件输出
func NewWowShopEmail() *WowShopEmail {
	return &WowShopEmail{}
}

// Process 处理获得的数据
func (thisObj *WowShopEmail) Process(items *page_items.PageItems, t com_interfaces.Task) {
	name := items.GetAll()["Name"]
	price := items.GetAll()["Price"]

	emailTool.SendMail([]string{"546247083@qq.com"}, "魔兽商城打折报告", name+price, true, nil)
}
