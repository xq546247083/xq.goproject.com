package rpcServer

//Priority 优先级
type Priority int8

const (
	//ConHighPriority 高优先级
	ConHighPriority Priority = 1 + iota

	//ConMediumPriority 中优先级
	ConMediumPriority

	//ConLowPriority 低优先级
	ConLowPriority
)
