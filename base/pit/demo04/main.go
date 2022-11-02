package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	字符串的长度
 */
func main()  {
	char := "♥"
	fmt.Println(len(char))    // 3

	//Go 的内建函数 len() 返回的是字符串的 byte 数量，而不是像 Python 中那样是计算 Unicode 字符数。
	//如果要得到字符串的字符数，可使用 "unicode/utf8" 包中的 RuneCountInString(str string) (n int)

	fmt.Println(utf8.RuneCountInString(char))

	char1 := "侯"
	fmt.Println(len(char1))
	fmt.Println(utf8.RuneCountInString(char1))

	fmt.Println("------------------")
	//注意： RuneCountInString 并不总是返回我们看到的字符数，因为有的字符会占用 2 个 rune：
	char3 := "é"
	fmt.Println(len(char3))    // 3 --- 事实上，这里并不是3，展示的2
	fmt.Println(utf8.RuneCountInString(char3))    // 2 --- 这里是1
	fmt.Println("cafe\u0301")    // café    // 法文的 cafe，实际上是两个 rune 的组合

}
