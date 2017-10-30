package spider

import (
	"strings"

	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page_items"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

// PipelineMysql mysql输出持久层
type PipelineMysql struct {
}

// NewPipelineMysql 新建mysql输出持久层
func NewPipelineMysql() *PipelineMysql {
	return &PipelineMysql{}
}

// Process 处理获得的数据
func (thisObj *PipelineMysql) Process(items *page_items.PageItems, t com_interfaces.Task) {
	name := items.GetAll()["name"]
	title := items.GetAll()["title"]
	source := strings.Replace(strings.Replace(items.GetAll()["source"], "↓", "", -1), " ", "", -1)
	content := items.GetAll()["content"]

	novel := model.NewNovel(name, title, source, content)
	dal.NovelDALObj.SaveInfo(novel, nil)
}
