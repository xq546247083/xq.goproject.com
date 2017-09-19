package webServer

import (
	"net/http"
	"strings"

	"xq.goproject.com/commonTool/configTool"
)

//读取配置数据
var (
	httpDir      = http.Dir(configTool.WebMainPath)
	indexPage    = configTool.IndexPage
	error404Page = configTool.Error404Page
	error500Page = configTool.Error500Page
)

//TimeFormat For parsing this time format, see ParseTime.
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

// The algorithm uses at most sniffLen bytes to make its decision.
const sniffLen = 512

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)
