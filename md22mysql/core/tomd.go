package core

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// 获取表列表
func getTableList(tempDataBaseName string) []*tableInfo {
	// 查询表信息
	queryTableInfo, err := mysqlConn.Query("SELECT TABLE_NAME ,table_comment FROM information_schema.`TABLES` WHERE table_schema = ? ORDER BY table_name", tempDataBaseName)
	if err != nil {
		log.Println("query table desc failed,err:", err.Error())
		os.Exit(1)
	}
	defer queryTableInfo.Close()

	// 赋值到对象
	var tableList []*tableInfo
	for queryTableInfo.Next() {
		tableInfoObj := new(tableInfo)
		if err := queryTableInfo.Scan(&tableInfoObj.Name, &tableInfoObj.Desc); err != nil {
			log.Println("scan table desc failed,err:", err.Error())
			os.Exit(1)
		}
		tableList = append(tableList, tableInfoObj)
	}

	return tableList
}

// 更新表的列信息
func updateTableColumnList(tables []*tableInfo) {
	// 获取表字符串
	var tableListStr string
	for _, tableObj := range tables {
		tableListStr = tableListStr + ",'" + tableObj.Name + "'"
	}
	tableListStr = strings.Trim(tableListStr, ",")

	// 查询表下面的列信息
	queryRowInfo, err := mysqlConn.Query(fmt.Sprintf("select TABLE_NAME,COLUMN_NAME,IS_NULLABLE, COLUMN_TYPE,COLUMN_DEFAULT,COLUMN_COMMENT,COLUMN_Key,Extra from information_schema.`COLUMNS` where TABLE_SCHEMA ='%s' and TABLE_NAME IN (%s)", DataBaseName, tableListStr))
	if err != nil {
		log.Println("select table info failed,err :", err.Error())
		os.Exit(1)
	}
	defer queryRowInfo.Close()

	// 更新表的列信息
	for queryRowInfo.Next() {
		var tableName string
		columnObj := new(columnInfo)
		if err := queryRowInfo.Scan(&tableName, &columnObj.Name, &columnObj.IsNullAble, &columnObj.Type, &columnObj.Default, &columnObj.Desc, &columnObj.Key, &columnObj.Extra); err != nil {
			log.Println("scan table failed,err:", err.Error())
			os.Exit(1)
		}
		for _, tableObj := range tables {
			if tableObj.Name == tableName {
				tableObj.ColumnList = append(tableObj.ColumnList, columnObj)
			}
		}
	}
}

// 转换为md
func doToMd(tables []*tableInfo) {
	// 创建文件
	filename := DataBaseName + "数据库.md"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Println("创建文件失败,错误:", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	// 字节流
	buf := bytes.Buffer{}

	// 循环表，写信息
	for _, tableObj := range tables {
		//表名
		buf.WriteString(fmt.Sprintf("### %s:%s\n", tableObj.Name, tableObj.Desc.String))
		//表内容
		buf.WriteString(columnTitle + "\n")
		buf.WriteString(columnSeparator + "\n")
		for _, j := range tableObj.ColumnList {
			// 获取默认字符串
			var defaultStr string
			if !j.Default.Valid {
				defaultStr = "null"
			} else {
				if isStringByColumnType(j.Type) {
					defaultStr = fmt.Sprintf("\"%s\"", j.Default.String)
				} else {
					defaultStr = fmt.Sprintf("%s", j.Default.String)
				}

			}

			// 获取类型
			var keyStr string
			if j.Key.Valid && j.Key.String == "PRI" {
				keyStr = "PK"
			}

			// 获取自动递增类型
			if j.Extra.Valid && strings.Contains(j.Extra.String, "auto_increment") {
				if keyStr == "" {
					keyStr = "AI"
				} else {
					keyStr = keyStr + ",AI"
				}
			}

			buf.WriteString(fmt.Sprintf("|%s|%s|%s|%s|%s|\"%s\"|\n", j.Name, j.Type, keyStr, j.IsNullAble, defaultStr, j.Desc.String))
		}

		buf.WriteString("\n<br>\n\n")
	}
	if _, err := file.Write(buf.Bytes()); err != nil {
		log.Println("file write fail,err:", err.Error())
		os.Exit(1)
	}
	fmt.Println("导出成功 ==>", filename)
}

// 转换md
func ToMd() {
	connectMysql()

	tableList := getTableList(DataBaseName)
	updateTableColumnList(tableList)
	doToMd(tableList)
}
