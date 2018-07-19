// 反射工具
package refelctTool

import (
	"reflect"
)

type reflectObj struct {
	Name  string
	Value reflect.Value
}

// 反射结构的所有方法
// structObject：对象
func GetMothodList(structObject interface{}) []*reflectObj {
	// 获取structObject对应的反射 Type 和 Value
	reflectValue := reflect.ValueOf(structObject)
	reflectType := reflect.TypeOf(structObject)

	result := make([]*reflectObj, 0, reflectType.NumMethod())

	// 获取structObject中返回值为responseObject的方法
	for i := 0; i < reflectType.NumMethod(); i++ {
		methodName := reflectType.Method(i).Name
		reflectObj := &reflectObj{
			Name:  methodName,
			Value: reflectValue.MethodByName(methodName),
		}

		result = append(result, reflectObj)
	}

	return result
}
