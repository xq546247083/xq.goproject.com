package webClient

// APIType API类型
type APIType string

var (
	// GetAllUser 获取所有用户的api
	GetAllUser APIType = "Func/SysUser/GetAllUser"

	// GetUser 获取用户的api
	GetUser APIType = "Func/SysUser/GetUser"
)
