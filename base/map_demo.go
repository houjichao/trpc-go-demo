package base

import "fmt"

/*
	map学习：

	map变量声明
	var map_variable map[key_data_type] value_data_type

	map初始化:
	使用make函数
	map_variable := make(map[key_data_type]value_data_type)

*/
func MapDemo() {
	var familyPersonMap map[string]string
	familyPersonMap = make(map[string]string)
	familyPersonMap["one"] = "侯吉超"
	familyPersonMap["two"] = "杨文菊"
	familyPersonMap["three"] = "侯怡菲"

	for name := range familyPersonMap {
		fmt.Println(name, " is ", familyPersonMap[name])
	}
}
