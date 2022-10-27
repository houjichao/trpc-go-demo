package base

import "fmt"

/*
	字符串学习
 */
func StrDemo()  {
	//1.如果字符串没有特殊字符，字符串的表现形式用双引号
	//2.如果字符串有特殊字符，字符串的表现形式用反引号
	var str string = `
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

	`
	fmt.Println(str)

	//字符串过长的拼接,+号要保留在上一行
	var str1 string = "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" +
		"abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" + "def" + "abc" + "def"
	fmt.Println(str1)
}