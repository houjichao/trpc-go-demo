package main

import "fmt"

/*
	map学习
*/
func main() {

	//创建方式1
	//定义map变量
	var map1 map[int]string
	//必须make初始化，才会分配空间
	//size 1可以省略，会自动扩容
	map1 = make(map[int]string, 1)
	map1[2001] = "张三"
	map1[2002] = "李四"
	map1[2003] = "王五"
	fmt.Println(map1)

	//方式2
	map2 := make(map[int]string, 2)
	map2[2001] = "张三"
	map2[2002] = "李四"
	map2[2003] = "王五"
	fmt.Println(map2)

	//方式3
	map3 := map[int]string{
		2001: "张三",
		2002: "李四",
	}
	map3[2003] = "王五"
	fmt.Println(map3)
}
