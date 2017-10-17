package webSocketServerObject

//ResponseObject Socket服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值(用SetResultStatus来设置响应结果的状态值)
	Status ResultStatus

	// 响应结果的状态值所对应的描述信息
	StatusValue string

	// 附加数据
	AttachData map[string]interface{}

	// 响应结果的数据
	Data interface{}
}

//NewResponseObject 创建一个新的服务器返回object
func NewResponseObject() *ResponseObject {

	return &ResponseObject{
		Status:      Success,
		StatusValue: Success.ToDescription(),
		AttachData:  make(map[string]interface{}),
		Data:        nil,
	}
}

//SetResultStatus 设置状态
func (responseObject *ResponseObject) SetResultStatus(rs ResultStatus) *ResponseObject {
	responseObject.Status = rs
	responseObject.StatusValue = rs.ToDescription()

	return responseObject
}
