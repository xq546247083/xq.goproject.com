package dal

import (
	"fmt"

	"xq.goproject.com/commonTools/stringTool"

	"xq.goproject.com/Vendor/github.com/jinzhu/gorm"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/logTool"
)

var (
	// DB 数据库
	DB *gorm.DB
)

func init() {
	dbConStr := configTool.DBConnection
	var err error
	if DB, err = gorm.Open("mysql", dbConStr); err != nil {
		logTool.Log(logTool.Error, fmt.Sprintf("连接数据库错误：%s", err))
		panic(fmt.Errorf("连接数据库错误：%s", err))
	}
}

func writeErrorLog(err error, errMsg string) {
	logTool.Log(logTool.Error, fmt.Sprintf("%s执行数据库操作错误.%serr:%s", errMsg, stringTool.GetNewLine(), err))
}
