package common

//OperateType 操作类型
type OperateType int

const (
	//AddOperate 加
	AddOperate OperateType = 1 + iota

	//ReduceOperate 减
	ReduceOperate
)
