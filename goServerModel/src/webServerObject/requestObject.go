package webServerObject

//RequestObject 客户端请求对象
type RequestObject struct {
	// 请求的方法名称
	MethodName string

	// 请求的参数数组
	Parameters []interface{}
}
