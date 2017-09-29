package configTool

import (
	"fmt"
	"strings"

	"xq.goproject.com/commonTools/typeTool"
	"xq.goproject.com/commonTools/xmlTool"
)

type XmlConfig struct {
	root *xmlTool.Node
}

// 从文件加载
// xmlFilePath:xml文件路径
// 返回值:
// error:错误信息
func (thisObj *XmlConfig) LoadFromFile(xmlFilePath string) error {
	if thisObj.root != nil {
		return fmt.Errorf("have loaded")
	}

	root, errMsg := xmlTool.LoadFromFile(xmlFilePath)
	if errMsg != nil {
		return errMsg
	}

	thisObj.root = root

	return nil
}

// 从node节点加载（会取其根节点）
// xmlRoot:xml节点
// 返回值:
// error:错误信息
func (thisObj *XmlConfig) LoadFromXmlNode(xmlRoot *xmlTool.Node) error {
	if thisObj.root != nil {
		return fmt.Errorf("have loaded")
	}

	if xmlRoot == nil {
		return fmt.Errorf("xmlRoot is nil")
	}

	thisObj.root = xmlRoot

	return nil
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// bool:结果
// error:错误信息
func (thisObj *XmlConfig) Bool(xpath string, attrName string) (bool, error) {
	val, errMsg := thisObj.getVal(xpath, attrName)
	if errMsg != nil {
		return false, errMsg
	}

	return typeTool.Bool(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// int:结果
// error:错误信息
func (thisObj *XmlConfig) Int(xpath string, attrName string) (int, error) {
	val, errMsg := thisObj.getVal(xpath, attrName)
	if errMsg != nil {
		return 0, errMsg
	}

	return typeTool.Int(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// int64:结果
// error:错误信息
func (thisObj *XmlConfig) Int64(xpath string, attrName string) (int64, error) {
	val, errMsg := thisObj.getVal(xpath, attrName)
	if errMsg != nil {
		return 0, errMsg
	}

	return typeTool.Int64(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// float64:结果
// error:错误信息
func (thisObj *XmlConfig) Float(xpath string, attrName string) (float64, error) {
	val, errMsg := thisObj.getVal(xpath, attrName)
	if errMsg != nil {
		return 0, errMsg
	}

	return typeTool.Float64(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// string:结果
// error:错误信息
func (thisObj *XmlConfig) String(xpath string, attrName string) (string, error) {
	return thisObj.getVal(xpath, attrName)
}

// 获取指定位置的节点
// xpath:xpath路径
// 返回值:
// []*xmlTool.Node：结果
func (thisObj *XmlConfig) Nodes(xpath string) []*xmlTool.Node {
	return thisObj.root.SelectElements(xpath)
}

// 获取指定位置的节点
// xpath:xpath路径
// 返回值:
// *xmlTool.Node：结果
func (thisObj *XmlConfig) Node(xpath string) *xmlTool.Node {
	return thisObj.root.SelectElement(xpath)
}

// 获取指定路径的之
// xpath:xpath路径
// attrName:要获取的属性值，如果为空，则返回内部文本
func (thisObj *XmlConfig) getVal(xpath string, attrName string) (string, error) {
	targetRoot := thisObj.root.SelectElement(xpath)
	if targetRoot == nil {
		return "", fmt.Errorf("no find target node:%v", xpath)
	}

	val := ""
	if attrName == "" {
		val = strings.TrimSpace(targetRoot.InnerText())
	} else {
		val = targetRoot.SelectAttr(attrName)
	}

	return val, nil
}

// NewXmlConfig 创建新的xml配置对象
func NewXmlConfig() *XmlConfig {
	return &XmlConfig{}
}
