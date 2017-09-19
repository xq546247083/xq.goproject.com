package rpcServerObject

//RequestObject 客户端请求对象
type RequestObject struct {
	// 请求的唯一标识，是需要通过截取请求数据前4位得到并进行手动赋值的
	ID int32

	// 请求的方法名称
	MethodName string

	// 请求的参数数组
	Parameters []interface{}
}
