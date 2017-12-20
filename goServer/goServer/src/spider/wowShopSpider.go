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
	header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	header.Set("Accept-Encoding", "gzip, deflate, br")
	header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	header.Set("Cookie", "web.id=CN-8385fe60-d079-46b4-af24-956b02fa9a4a; __utma=124133273.462908016.1513412632.1513412821.1513738956.2; __utmz=124133273.1513738956.2.2.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; XSRF-TOKEN=10bf4390-29bd-497c-aa54-7ffd9f100c6c; loc=zh-cn; _ga=GA1.4.462908016.1513412632; _gid=GA1.4.1119770171.1513738955; _gat_UA-50249600-1=1; _ga=GA1.3.462908016.1513412632; _gid=GA1.3.1119770171.1513738955")
	header.Set("Host", "shop.battlenet.com.cn")
	header.Set("Upgrade-Insecure-Requests", "1")
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	req := request.NewRequest("https://shop.battlenet.com.cn/zh-cn/product/world-of-warcraft-service-faction-change", "html", "WowShopSpider", "GET", "", header, nil, nil, nil)
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
