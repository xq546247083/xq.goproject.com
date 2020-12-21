package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

// 列title
const columnTitle = "|字段|类型|主键|是否为空|默认|注释|"

// 列分割线
const columnSeparator = "| :--: | :--: | :--: | :--: | :--: | :--: |"

var (
	// 数据库用户名
	MysqlUserName string

	// 数据库密码
	MysqlPwd string

	// 数据库地址
	MysqlNetAddress string

	// 数据库名字
	DataBaseName string

	// toMysql的MD文件路径
	ToMysqlMdFilePath string

	// 数据库连接
	mysqlConn *sql.DB
)

// 连接mysql
func connectMysql() {
	sqlString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", MysqlUserName, MysqlPwd, MysqlNetAddress, DataBaseName)

	var err error
	mysqlConn, err = sql.Open("mysql", sqlString)
	if err != nil {
		log.Println("数据库连接失败,错误:", err.Error())
		os.Exit(1)
	}
}

// 通过行，获取默认值是否是字符串
func isStringByColumnType(columnType string) bool {
	columnType = strings.ToLower(columnType)
	if strings.Contains(columnType, "char") ||
		strings.Contains(columnType, "text") ||
		strings.Contains(columnType, "date") ||
		strings.Contains(columnType, "string") {
		return true
	}

	return false
}
