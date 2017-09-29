package query

import "xq.goproject.com/commonTools/xmlTool/gxpath/xpath"

type Iterator interface {
	Current() xpath.NodeNavigator
}
