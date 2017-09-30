package enum

// BlogStatusEnum 博客类型枚举
type BlogStatusEnum byte

const (
	// BlogDraft 加
	BlogDraft BlogStatusEnum = iota

	// BlogCommon 减
	BlogCommon

	// BlogDelete 删除
	BlogDelete

	// BlogRemove 彻底删除
	BlogRemove
)
