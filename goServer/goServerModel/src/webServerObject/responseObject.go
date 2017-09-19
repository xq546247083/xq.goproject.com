package webServerObject

//ResponseObject Socket服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值(用SetResultStatus来设置响应结果的状态值)
	Status ResultStatus

	// 响应结果的状态值所对应的描述信息
	Message string

	// 响应结果的数据
	Data interface{}
}

//NewResponseObject 创建一个新的服务器返回object
func NewResponseObject() *ResponseObject {
	return &ResponseObject{
		Status:  Success,
		Message: Success.ToString(),
		Data:    nil,
	}
}

//SetResultStatus 设置状态
func (responseObject *ResponseObject) SetResultStatus(rs ResultStatus) *ResponseObject {
	responseObject.Status = rs
	responseObject.Message = rs.ToString()

	return responseObject
}

//SetData 设置数据
func (responseObject *ResponseObject) SetData(data interface{}) *ResponseObject {
	responseObject.Data = data

	return responseObject
}
