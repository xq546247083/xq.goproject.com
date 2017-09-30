package webServerObject

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"xq.goproject.com/commonTools/typeTool"

	"xq.goproject.com/commonTools/logTool"
)

// RequestObject http请求
type RequestObject struct {
	// HTTPRequest 请求对象
	HTTPRequest *http.Request

	// requestInfo 请求信息
	requestInfo map[string]interface{}
}

// NewRequestObject 新建http请求
func NewRequestObject(_request *http.Request) *RequestObject {
	return &RequestObject{
		HTTPRequest: _request,
	}
}

// 获取请求的字段
func getRequsetByte(requestObj *RequestObject) ([]byte, error) {
	defer func() {
		requestObj.HTTPRequest.Body.Close()
	}()

	data, err := ioutil.ReadAll(requestObj.HTTPRequest.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetObj 获取对象
func (thisObj *RequestObject) GetObj(name string) (interface{}, error) {
	// 如果没有序列化过请求，则序列化请求
	if len(thisObj.requestInfo) <= 0 {
		data, err := getRequsetByte(thisObj)
		if data == nil || err != nil {
			return nil, errors.New("RequestBytes为空")
		}

		// 反序列化
		if err := json.Unmarshal(data, &thisObj.requestInfo); err != nil {
			logTool.Log(logTool.Error, fmt.Sprintf("反序列化失败，字符串为：%s.err:%s", string(data), err))
			return nil, err
		}
	}

	strObj, exists := thisObj.requestInfo[name]
	if !exists {
		return nil, fmt.Errorf("不存在%s的字段", name)
	}

	return strObj, nil
}

// GetStringVal 获取请求的值
func (thisObj *RequestObject) GetStringVal(name string) (string, error) {
	obj, err := thisObj.GetObj(name)
	if err != nil {
		return "", err
	}

	str, err := typeTool.String(obj)
	if err != nil {
		return "", err
	}

	return str, nil
}

// GetIntVal 获取请求的值
func (thisObj *RequestObject) GetIntVal(name string) (int, error) {
	obj, err := thisObj.GetObj(name)
	if err != nil {
		return -1, nil
	}

	intVal, err := typeTool.Int(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// GetInt32Val 获取请求的值
func (thisObj *RequestObject) GetInt32Val(name string) (int32, error) {
	obj, err := thisObj.GetObj(name)
	if err != nil {
		return -1, nil
	}

	intVal, err := typeTool.Int32(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// GetInt64Val 获取请求的值
func (thisObj *RequestObject) GetInt64Val(name string) (int64, error) {
	obj, err := thisObj.GetObj(name)
	if err != nil {
		return -1, nil
	}

	intVal, err := typeTool.Int64(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// GetFloat64Val 获取请求的值
func (thisObj *RequestObject) GetFloat64Val(name string) (float64, error) {
	obj, err := thisObj.GetObj(name)
	if err != nil {
		return -1, nil
	}

	intVal, err := typeTool.Float64(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}
