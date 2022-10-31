package main

import "fmt"

/*
	切片的循环和注意事项
*/
func main() {
	var intarr = [6]int{3, 6, 9, 1, 4, 7}
	var slice1 = intarr[1:5]

	for k, v := range slice1 {
		fmt.Println(k, v)
	}

	//切片可以动态增长
	slice2 := append(slice1, 88, 50)
	fmt.Println(slice2)
	//底层原理:
	//1.底层追加元素的时候对老数组进行扩容，老数组扩容为新数组
	//2.创建一个新数组，将老数组中的原数据复制到新数组中，在新数组中追加88，50
	//3.slice2 底层数组的指向的是新数组
	//4.往往我们在使用追加的时候想要的是给原数组追加
	slice1 = append(slice1, 88, 59)
	fmt.Println(slice1)
	//5.底层的新数组，不能直接维护，要通过切片间接维护

	//切片可以追加切片
	slice3 := []int{99, 101}
	slice1 = append(slice1, slice3...)
	fmt.Println(slice1)

	//切片拷贝
	var a = []int{1, 4, 7, 3, 6, 9}
	var b = make([]int, 10)
	copy(b, a) //将a拷贝到b
	fmt.Println(a)
	fmt.Println(b)
}
