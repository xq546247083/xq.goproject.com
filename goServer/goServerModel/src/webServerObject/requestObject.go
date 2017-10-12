package webServerObject

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/commonTools/typeTool"
)

// RequestObject http请求
type RequestObject struct {
	// HTTPRequest 请求对象
	HTTPRequest *http.Request

	// requestInfo 请求信息
	requestInfo map[string]interface{}

	// data 请求数据
	data []interface{}

	// isRead 是否读取
	isRead bool
}

// NewRequestObject 新建http请求
func NewRequestObject(_request *http.Request) *RequestObject {
	return &RequestObject{
		HTTPRequest: _request,
		isRead:      false,
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

//--------------------------------------下面获取val数据------------------------------------------------------

// GetValStr 获取请求的数据
func (thisObj *RequestObject) GetValStr() (string, error) {
	// 如果没有序列化过请求，则序列化请求
	if !thisObj.isRead {
		thisObj.isRead = true

		dataTemp, err := getRequsetByte(thisObj)
		if dataTemp == nil || err != nil {
			return "", errors.New("RequestBytes为空")
		}

		// 反序列化
		if err := json.Unmarshal(dataTemp, &thisObj.requestInfo); err != nil {
			logTool.Log(logTool.Error, fmt.Sprintf("反序列化失败，字符串为：%s.err:%s", string(dataTemp), err))
			return "", err
		}
	}

	dataByte, err := json.Marshal(thisObj.requestInfo)
	if err != nil {
		return "", err
	}

	return string(dataByte), nil
}

// getObjVal 获取对象
func (thisObj *RequestObject) getObjVal(name string) (interface{}, error) {
	// 如果没有序列化过请求，则序列化请求
	if !thisObj.isRead {
		thisObj.isRead = true

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

// GetStringVal 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetStringVal(name string) (string, error) {
	obj, err := thisObj.getObjVal(name)
	if err != nil {
		return "", err
	}

	str, err := typeTool.String(obj)
	if err != nil {
		return "", err
	}

	return str, nil
}

// GetIntVal 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetIntVal(name string) (int, error) {
	obj, err := thisObj.getObjVal(name)
	if err != nil {
		return -1, err
	}

	intVal, err := typeTool.Int(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// GetInt32Val 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetInt32Val(name string) (int32, error) {
	obj, err := thisObj.getObjVal(name)
	if err != nil {
		return -1, err
	}

	intVal, err := typeTool.Int32(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// GetInt64Val 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetInt64Val(name string) (int64, error) {
	obj, err := thisObj.getObjVal(name)
	if err != nil {
		return -1, err
	}

	intVal, err := typeTool.Int64(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// GetFloat64Val 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetFloat64Val(name string) (float64, error) {
	obj, err := thisObj.getObjVal(name)
	if err != nil {
		return -1, err
	}

	intVal, err := typeTool.Float64(obj)
	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// GetSliceObjectVal 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetSliceObjectVal(name string) ([]interface{}, error) {
	obj, err := thisObj.getObjVal(name)
	if err != nil {
		return nil, err
	}

	varArray := obj.([]interface{})

	return varArray, nil
}

//--------------------------------------下面获取data数据------------------------------------------------------

// GetData 获取请求的值(Data：请求参数值数据)
func (thisObj *RequestObject) GetData() ([]interface{}, error) {
	// 如果转换数据，则转换数据
	if len(thisObj.data) <= 0 {
		var err error
		thisObj.data, err = thisObj.GetSliceObjectVal("Data")
		if err != nil {
			return nil, err
		}

		if len(thisObj.data) <= 0 {
			return nil, errors.New(ParamIsEmpty.ToDescription())
		}
	}

	return thisObj.data, nil
}

// getObjectData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) getObjectData(num int32) (interface{}, error) {
	_, err := thisObj.GetData()
	if err != nil {
		return "", err
	}

	if len(thisObj.data) < int(num) {
		return "", errors.New(APIDataError.ToDescription())
	}

	return thisObj.data[num-1], nil
}

// GetStringData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetStringData(num int32) (string, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return "", err
	}

	str, err := typeTool.String(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return "", err
	}

	return str, nil
}

// GetIntData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetIntData(num int32) (int, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	intVal, err := typeTool.Int(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return intVal, nil
}

// GetInt32Data 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetInt32Data(num int32) (int32, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	int32Val, err := typeTool.Int32(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return int32Val, nil
}

// GetInt64Data 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetInt64Data(num int32) (int64, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	int64Val, err := typeTool.Int64(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return int64Val, nil
}

// GetFloat64Data 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetFloat64Data(num int32) (float64, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	float64CVal, err := typeTool.Float64(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return float64CVal, nil
}

// GetBoolData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetBoolData(num int32) (bool, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return false, err
	}

	boolVal, err := typeTool.Bool(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.HTTPRequest.RequestURI, stringTool.GetNewLine(), thisObj.data, stringTool.GetNewLine(), num), err.Error())
		return false, err
	}

	return boolVal, nil
}
