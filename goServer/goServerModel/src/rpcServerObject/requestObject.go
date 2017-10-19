package rpcServerObject

import (
	"errors"
	"fmt"

	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/commonTools/typeTool"
)

//RequestObject 客户端请求对象
type RequestObject struct {
	// 请求的方法名称
	MethodName string

	// RequestInfo 请求信息
	RequestInfo map[string]interface{}
}

// NewRequestObject 新建RpcHttp请求
func NewRequestObject() *RequestObject {
	return &RequestObject{
		RequestInfo: make(map[string]interface{}),
	}
}

//--------------------------------------下面获取val数据------------------------------------------------------

// GetObjVal 获取对象
func (thisObj *RequestObject) GetObjVal(name string) (interface{}, error) {
	strObj, exists := thisObj.RequestInfo[name]
	if !exists {
		return nil, fmt.Errorf("不存在%s的字段", name)
	}

	return strObj, nil
}

// GetStringVal 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetStringVal(name string) (string, error) {
	obj, err := thisObj.GetObjVal(name)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return "", err
	}

	str, err := typeTool.String(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return "", err
	}

	return str, nil
}

// GetIntVal 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetIntVal(name string) (int, error) {
	obj, err := thisObj.GetObjVal(name)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	intVal, err := typeTool.Int(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	return intVal, nil
}

// GetInt32Val 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetInt32Val(name string) (int32, error) {
	obj, err := thisObj.GetObjVal(name)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	intVal, err := typeTool.Int32(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	return intVal, nil
}

// GetInt64Val 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetInt64Val(name string) (int64, error) {
	obj, err := thisObj.GetObjVal(name)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	intVal, err := typeTool.Int64(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	return intVal, nil
}

// GetFloat64Val 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetFloat64Val(name string) (float64, error) {
	obj, err := thisObj.GetObjVal(name)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	intVal, err := typeTool.Float64(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取【%s】失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), name), err.Error())
		return -1, err
	}

	return intVal, nil
}

// GetSliceObjectVal 获取请求的值(Val:非参数值)
func (thisObj *RequestObject) GetSliceObjectVal(name string) ([]interface{}, error) {
	obj, err := thisObj.GetObjVal(name)
	if err != nil {
		return nil, err
	}

	varArray := obj.([]interface{})

	return varArray, nil
}

//--------------------------------------下面获取data数据------------------------------------------------------

// GetData 获取请求的值(Data：请求参数值数据)
func (thisObj *RequestObject) GetData() ([]interface{}, error) {
	return thisObj.GetSliceObjectVal("Data")
}

// getObjectData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) getObjectData(num int32) (interface{}, error) {
	data, err := thisObj.GetData()
	if err != nil {
		return "", err
	}

	if len(data) < int(num) {
		return "", errors.New(APIDataError.ToDescription())
	}

	return data[num-1], nil
}

// GetStringData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetStringData(num int32) (string, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return "", err
	}

	str, err := typeTool.String(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return "", err
	}

	return str, nil
}

// GetIntData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetIntData(num int32) (int, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	intVal, err := typeTool.Int(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return intVal, nil
}

// GetInt32Data 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetInt32Data(num int32) (int32, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	int32Val, err := typeTool.Int32(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return int32Val, nil
}

// GetInt64Data 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetInt64Data(num int32) (int64, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	int64Val, err := typeTool.Int64(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return int64Val, nil
}

// GetFloat64Data 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetFloat64Data(num int32) (float64, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	float64CVal, err := typeTool.Float64(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return -1, err
	}

	return float64CVal, nil
}

// GetBoolData 获取请求的值(Data：请求参数值数据)
//参数值：num（从1开始）
func (thisObj *RequestObject) GetBoolData(num int32) (bool, error) {
	obj, err := thisObj.getObjectData(num)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return false, err
	}

	boolVal, err := typeTool.Bool(obj)
	if err != nil {
		logTool.LogError(fmt.Sprintf("请求地址：%s %s请求数据：%s %s获取第%d个参数失败", thisObj.MethodName, stringTool.GetNewLine(), thisObj.RequestInfo, stringTool.GetNewLine(), num), err.Error())
		return false, err
	}

	return boolVal, nil
}
