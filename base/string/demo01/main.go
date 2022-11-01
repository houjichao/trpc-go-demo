package main

import (
	"fmt"
	"strings"
)

/*
Go语言的字符有以下两种：
一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。
*/
func main() {

	toUpper := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			r = r - 32
		}
		//r = r + ('A' - 'a')
		return r
	}
	input := "Hello, Welcome to GeeksforGeeks"

	result := strings.Map(toUpper, input)
	fmt.Println(result)
}
