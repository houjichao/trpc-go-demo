package main

import "fmt"

/*
	数组学习
*/
func main() {
	//定义语法
	// var 数组名 [size] 数据类型
	var scores [5]int
	//将成绩存入数组：（循环 + 终端输入）
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第%d个学生的成绩", i+1)
		fmt.Scanln(&scores[i])
	}
	fmt.Println("--------------")
	//数组遍历
	//普通循环
	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩为：%d\n", i+1, scores[i])
	}
	//for-range
	for key, value := range scores {
		fmt.Printf("第%d个学生的成绩为：%d\n", key+1, value)
	}
	//for-range 忽略key或者value，可以用下划线
	for _, value := range scores {
		fmt.Printf("学生的成绩为：%d\n", value)
	}
}
