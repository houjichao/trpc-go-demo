package main

import "fmt"

/*
	数组的几个注意事项
*/
func main() {
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)
	//注意1：长度属于类型的一部分
	fmt.Printf("数组的类型为:%T \n", arr1)

	test(arr1)
	//注意2：go中数组属值类型，在默认情况下是值传递，因此会进行值拷贝
	fmt.Println(arr1) //[1 2 3]

	//注意3：如想在其它函数中，去修改原来的数组，可以使用使用传递（指针方式）
	test1(&arr1)
	fmt.Println(arr1)
}

func test(arr [3]int) {
	arr[0] = 7
}

func test1(arr *[3]int) {
	//(*arr)[0] = 7
	arr[0] = 7
}
