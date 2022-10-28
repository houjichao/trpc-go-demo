package main

import "fmt"

func main()  {

	var age int

	fmt.Println("请录入学生的年龄：")
	fmt.Scanln(&age)
	fmt.Println("学生年龄为：",age)

}
