package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	// 数据库用户名
	MysqlUserName string

	// 数据库密码
	MysqlPwd string

	// 数据库地址
	MysqlNetAddress string

	// 数据库名字
	DataBaseName string

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
