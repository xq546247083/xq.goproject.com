package main

import (
	"fmt"

	"xq.goproject.com/commonTools/dataSyncTool/mysqlSync"
	"xq.goproject.com/commonTools/logTool"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var _ = mysql.DeregisterLocalFile

var (
	connectionString = "root:123456@tcp(localhost:3306)/webserver?charset=utf8&parseTime=true&loc=Local&timeout=60s"
	maxOpenConns     = 10
	maxIdleConns     = 10

	syncFileSize = 1024 * 1024
)

var (
	// 数据库对象
	dbObj *gorm.DB

	// 同步管理对象
	syncMgr *mysqlSync.SyncMgr
)

func init() {
	// 初始化数据库连接
	dbObj = initMysql()

	// 构造同步管理对象
	syncMgr = mysqlSync.NewSyncMgr("Sync", syncFileSize, 1, dbObj.DB())
}

// 初始化Mysql
func initMysql() *gorm.DB {
	dbObj, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(fmt.Errorf("初始化数据库:%s失败，错误信息为：%s", connectionString, err))
	}
	logTool.LogDebug("连接mysql:%s成功", connectionString)

	if maxOpenConns > 0 && maxIdleConns > 0 {
		dbObj.DB().SetMaxOpenConns(maxOpenConns)
		dbObj.DB().SetMaxIdleConns(maxIdleConns)
	}

	return dbObj
}

// 注册同步对象
func registerSyncObj(identifier string) {
	syncMgr.RegisterSyncObj(identifier)
}

// 保存sql数据
func save(identifier string, command string) {
	syncMgr.Save(identifier, command)
}
