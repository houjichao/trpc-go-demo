package main

import (
	"flag"
	"fmt"
)

func main() {

	var str string

	fmt.Println("开始解析")

	flag.StringVar(&str, "str", "", "parse string")

	//解析参数
	flag.Parse()
	//输出
	fmt.Println(str)

	//go run main.go -str="hello go"
}
