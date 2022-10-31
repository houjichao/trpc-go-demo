package main

import "fmt"

/*
	数组学习
*/
func main() {
	//定义语法
	// var 数组名 [size] 数据类型
	var scores [5]int
	scores[0] = 95
	scores[1] = 90
	scores[2] = 83
	scores[3] = 78
	scores[4] = 62

	//注意，这里即使第五个没有赋值，scores的长度依然为5
	fmt.Println(len(scores))

	//定义变长数组
	var names = [...]string{"hjc", "jack", "mary"}
	fmt.Println(len(names))

	//求和
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	//求平均
	avg := sum / len(scores)

	fmt.Printf("成绩的总分为:%v,成绩的平均数为%v \n", sum, avg)

}
