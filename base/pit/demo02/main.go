package main

import "fmt"

/*
访问 map 中不存在的 key
Go 则会返回元素对应数据类型的零值，比如 nil、'' 、false 和 0，取值操作总有值返回，故不能通过取出来的值来判断 key 是不是在 map 中。
*/
func main() {

	//错误示例
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if v := x["two"]; v == "" {
		fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
	}

	//正确示例
	y := map[string]string{"one": "1", "two": "2", "three": ""}
	if _, ok := y["three"]; !ok {
		fmt.Println("key three is no entry")
	}
}
