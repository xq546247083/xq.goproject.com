package dal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/logTool"
)

var (
	dbObj *gorm.DB
)

func init() {
	dbConStr := configTool.DBConnection
	var err error
	if dbObj, err = gorm.Open("mysql", dbConStr); err != nil {
		logTool.Log(logTool.Error, fmt.Sprintf("连接数据库错误：%s", err))
		panic(fmt.Errorf("连接数据库错误：%s", err))
	}
}

//GetDB 获取数据库
func GetDB() *gorm.DB {
	return dbObj
}
