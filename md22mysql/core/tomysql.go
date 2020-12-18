package core

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 转换为mysql语句
func ToMysql() {
	getTableInfoFromFile()
}

func getTableInfoFromFile() []*tableInfo {
	// 拼凑文件路径
	filePath, err := filepath.Abs(ToMysqlMdFilePath)
	if err != nil {
		log.Println("获取路径失败,错误:", err.Error())
		os.Exit(1)
	}

	// 读取文件
	mdByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("读取文件失败,错误:", err.Error())
		os.Exit(1)
	}

	mdStr := string(mdByte)
	strLineList := strings.Split(mdStr, "\n")

	var tableInfoList []*tableInfo
	var currentTableObj *tableInfo

	// 循环，组装sql
	for _, strLine := range strLineList {
		strLine = strings.TrimSpace(strLine)
		// 读取表信息
		if len(strLine) > 3 && strLine[0:3] == "###" {
			// 获取表的信息
			tableInfoSplitList := strings.Split(strLine[4:], ":")
			currentTableObj = new(tableInfo)
			if len(tableInfoSplitList) >= 2 {
				currentTableObj.Name = tableInfoSplitList[0]
				currentTableObj.Desc = sql.NullString{String: tableInfoSplitList[1], Valid: true}
			}

			// 添加表对象
			tableInfoList = append(tableInfoList, currentTableObj)
		} else if strLine == columnTitle || strLine == columnSeparator {
			continue
		} else if len(strLine) > 1 && strLine[0:1] == "|" && currentTableObj != nil {
			// 读取行信息
			columnInfoSplitList := strings.Split(strLine[1:len(strLine)-1], "|")
			if len(columnInfoSplitList) >= 6 {
				// 读取类型
				keyObj := sql.NullString{Valid: true}
				if columnInfoSplitList[2] == "PK" {
					keyObj.String = "PRI"
				}

				// 读取默认值
				var defaultObj sql.NullString
				if columnInfoSplitList[4] == "null" || columnInfoSplitList[4] == "" {
					defaultObj.Valid = false
				} else {
					defaultObj.Valid = true
					defaultObj.String = strings.Trim(columnInfoSplitList[4], "\"")
				}

				// 新建行对象
				columnObj := &columnInfo{
					Name:       columnInfoSplitList[0],
					Type:       columnInfoSplitList[1],
					Key:        keyObj,
					IsNullAble: columnInfoSplitList[3],
					Default:    defaultObj,
					Desc:       sql.NullString{Valid: true, String: strings.Trim(columnInfoSplitList[5], "\"")},
				}

				currentTableObj.ColumnList = append(currentTableObj.ColumnList, columnObj)
			}
		} else {
			currentTableObj = nil
		}
	}

	return tableInfoList
}
