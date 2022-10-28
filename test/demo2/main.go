package main

import "fmt"

/*
	学习循环
*/
func main() {

	var str string = "hello golang"
	//1.普通for循环
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	//2.for range
	var str1 string = "hello golang你好"
	for i, value := range str1 {
		fmt.Printf("索引为：%d,具体的值为：%c \n", i, value)
	}

}
