package configTool

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
)

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
