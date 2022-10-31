package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	学习字符串的一些系统内置函数
*/

func main() {
	//统计字符串的长度
	str := "golang你好" //go中，汉字是utf-8字符集，一个汉字3个字节
	fmt.Println("str长度：", len(str))

	//
	type str1 string
	const constName1 = "jack"
	var temp string = "hello"
	temp = constName1
	fmt.Println(temp)

	//字符循环
	//方式1：for range
	for i, v := range str {
		fmt.Printf("str的下标为%d,值为%c\n", i, v)
	}
	//方式2：切片 利用 r:=[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \n", r[i])
	}

	//整数转字符串
	str2 := strconv.Itoa(6887)
	fmt.Printf("%T \n", str2)

	//字符串转整数
	num1, _ := strconv.Atoi("122")
	fmt.Println(num1)

	//查找子串是否在指定的字符串中
	containsFlag := strings.Contains("javaandgolang", "go")
	fmt.Println(containsFlag)

	//统计字符串有几个指定的子串
	count := strings.Count("javagoandgolang", "go")
	fmt.Println(count)

	//不区分大小写的字符串比较
	equalFlag := strings.EqualFold("go", "Go")
	fmt.Println(equalFlag)

	//区分大小写的比较
	fmt.Println("hello" == "Hello")

	//返回子串在字符串第一次出现的索引值，没有返回-1
	index := strings.Index("javagoandgolang", "ga")
	fmt.Println(index)

	//字符串替换
	//第三个参数为几就代表替换几个，-1代表全部替换
	str3 := strings.Replace("goandjavagogo", "go", "golang", -1)
	fmt.Println(str3)

	//字符串切割
	strs := strings.Split("go-py-java","-")
	fmt.Println(len(strs))
	fmt.Println(strs)

	//大小写
	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("go"))

	//去空格
	fmt.Println(strings.TrimSpace("  go  java   "))
	//去指定字符
	fmt.Println(strings.Trim("~golang~","~"))
	fmt.Println(strings.TrimLeft("~golang~","~"))
	fmt.Println(strings.TrimRight("~golang~","~"))

	//判断是否以指定的字符串开头
	fmt.Println(strings.HasPrefix("http://www.baidu.com","http"))

	//判断是否以指定的字符串结尾
	fmt.Println(strings.HasSuffix("http://www.baidu.com","com"))
}
