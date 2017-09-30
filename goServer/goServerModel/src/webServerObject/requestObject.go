package webServerObject

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"xq.goproject.com/commonTools/logTool"
)

// RequestObject http请求
type RequestObject struct {
	// HTTPRequest 请求对象
	HTTPRequest *http.Request
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

// Unmarshal 反序列化
func (thisObj *RequestObject) Unmarshal(obj interface{}) error {
	data, err := getRequsetByte(thisObj)
	if data == nil || err != nil {
		return errors.New("RequestBytes为空")
	}

	//反序列化
	if err := json.Unmarshal(data, &obj); err != nil {
		logTool.Log(logTool.Error, fmt.Sprintf("反序列化失败，字符串为：%s.err:%s", string(data), err))
		return err
	}

	return nil
}
