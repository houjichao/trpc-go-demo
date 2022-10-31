package main

import "fmt"

/*
	二维数组学习
*/
func main() {
	var arr1 [2][3]int16
	fmt.Println(arr1)

	//二维数组内存分析
	fmt.Printf("arr的地址是:%p \n", &arr1)
	fmt.Printf("arr[0]的地址是:%p \n", &arr1[0])
	fmt.Printf("arr[0][0]的地址是:%p \n", &arr1[0][0])

	fmt.Printf("arr[1]的地址是:%p \n", &arr1[1])
	fmt.Printf("arr[1][0]的地址是:%p \n", &arr1[1][0])

	//赋值
	arr1[0][1] = 47
	fmt.Println(arr1)

	//初始化操作
	var arr2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)

	//二维数组循环
	var arr3 = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	//方式1:普通for循环
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Print(arr3[i][j], "\t")
		}
		fmt.Println()
	}

	//方式2：for-range循环
	for k1, v1 := range arr3 {
		for k2, v2 := range v1 {
			fmt.Printf("arr[%v][%v]=%v \t", k1, k2, v2)
		}
		fmt.Println()
	}
}
