package configTool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	//LogPath 日志路径
	LogPath = "Log"

	//WebMainPath 网站路径
	WebMainPath = "WebMain"

	//NecessaryField 服务器启动必要的配置字段
	NecessaryField string

	//RPCListenAddress 监听地址
	RPCListenAddress string

	//WebListenAddress 监听地址
	WebListenAddress string

	//IsDebug 是否测试模式
	IsDebug = false

	//DBConnection 数据库地址
	DBConnection string

	//LogInfoFlag 是否记录消息
	LogInfoFlag = true

	//LogDebugFlag 是否记录Debug消息
	LogDebugFlag = false

	//LogWarnFlag 是否记录警告消息
	LogWarnFlag = false

	//LogErrorFlag 是否记录错误消息
	LogErrorFlag = true

	//LogFatalFlag 是否记录致命错误消息消息
	LogFatalFlag = true

	//读取的配置
	config map[string]interface{}

	//错误
	err error
)

func init() {
	config, err = ReadConfig("config.ini")
	checkError(err, true)

	NecessaryField, err = ReadStringJSONValue(config, "NecessaryField")
	checkError(err, true)

	filedList := strings.Split(NecessaryField, ",")

	//读取路径
	LogPath, err = ReadStringJSONValue(config, "LogPath")
	checkError(err, isExist(filedList, "LogPath"))

	WebMainPath, err = ReadStringJSONValue(config, "WebMainPath")
	checkError(err, isExist(filedList, "WebMainPath"))

	RPCListenAddress, err = ReadStringJSONValue(config, "RPCListenAddress")
	checkError(err, isExist(filedList, "RPCListenAddress"))

	WebListenAddress, err = ReadStringJSONValue(config, "WebListenAddress")
	checkError(err, isExist(filedList, "WebListenAddress"))

	DBConnection, err = ReadStringJSONValue(config, "DBConnection")
	checkError(err, isExist(filedList, "DBConnection"))

	IsDebug, err = ReadBoolJSONValue(config, "IsDebug")
	checkError(err, isExist(filedList, "IsDebug"))

	//读取是否写日志
	LogInfoFlag, err = ReadBoolJSONValue(config, "LogInfoFlag")
	checkError(err, isExist(filedList, "LogInfoFlag"))

	LogDebugFlag, err = ReadBoolJSONValue(config, "LogDebugFlag")
	checkError(err, isExist(filedList, "LogDebugFlag"))

	LogWarnFlag, err = ReadBoolJSONValue(config, "LogWarnFlag")
	checkError(err, isExist(filedList, "LogWarnFlag"))

	LogErrorFlag, err = ReadBoolJSONValue(config, "LogErrorFlag")
	checkError(err, isExist(filedList, "LogErrorFlag"))

	LogFatalFlag, err = ReadBoolJSONValue(config, "LogFatalFlag")
	checkError(err, isExist(filedList, "LogFatalFlag"))
}

//ReadConfig 读取配置文件
func ReadConfig(path string) (map[string]interface{}, error) {
	readByte, error := ioutil.ReadFile(path)
	if error != nil {
		return nil, fmt.Errorf("读取配置文件(%s)出错:%s", path, error)
	}

	config := make(map[string]interface{})
	if error = json.Unmarshal(readByte, &config); error != nil {
		return nil, fmt.Errorf("序列化配置文件错:%s", error)
	}

	return config, nil
}

//ReadIntJSONValue 从config配置中获取int类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
// 错误对象
func ReadIntJSONValue(config map[string]interface{}, configName string) (int, error) {
	configValue, ok := config[configName]
	if !ok {
		return 0, fmt.Errorf("不存在名为%s的配置或配置为空", configName)
	}
	configValueFloat, ok := configValue.(float64)
	if !ok {
		return 0, fmt.Errorf("%s必须为int型", configName)
	}

	return int(configValueFloat), nil
}

//ReadStringJSONValue 从config配置中获取string类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
// 错误对象
func ReadStringJSONValue(config map[string]interface{}, configName string) (string, error) {
	configValue, ok := config[configName]
	if !ok {
		return "", fmt.Errorf("不存在名为%s的配置或配置为空", configName)
	}
	configValueString, ok := configValue.(string)
	if !ok {
		return "", fmt.Errorf("%s必须为string型", configName)
	}

	return configValueString, nil
}

//ReadBoolJSONValue 从config配置中获取string类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
// 错误对象
func ReadBoolJSONValue(config map[string]interface{}, configName string) (bool, error) {
	configValue, ok := config[configName]
	if !ok {
		return false, fmt.Errorf("不存在名为%s的配置或配置为空", configName)
	}
	configValueBool, ok := configValue.(bool)
	if !ok {
		return false, fmt.Errorf("%s必须为bool型", configName)
	}

	return configValueBool, nil
}

//是否存在
func isExist(listObj []string, obj string) bool {
	for _, item := range listObj {
		if item == obj {
			return true
		}
	}

	return false
}

//checkError 抛出错误
//是否抛出错误
func checkError(err error, ifThrowError bool) {
	if err != nil && ifThrowError {
		panic(err)
	}
}
