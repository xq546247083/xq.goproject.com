package core

import "database/sql"

// 列信息
type columnInfo struct {
	// 列名
	Name string

	// 列类型
	Type string

	// 是否可空
	IsNullAble string

	// 行默认值
	Default sql.NullString

	// 行描述
	Desc sql.NullString

	// 列上的索引类型 主键-->PRI  | 唯一索引 -->UNI  一般索引 -->MUL
	Key sql.NullString
}

// 表信息
type tableInfo struct {
	// 表名
	Name string

	// 描述
	Desc sql.NullString

	// 列信息
	ColumnList []*columnInfo
}
