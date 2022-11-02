package main

import "fmt"

//从 slice 中重新切出新 slice 时，新 slice 会引用原 slice 的底层数组。
//如果跳了这个坑，程序可能会分配大量的临时 slice 来指向原底层数组的部分数据，将导致难以预料的内存使用。
func get() []int {
	raw := make([]int, 10000)
	fmt.Printf("%T\n", raw)
	fmt.Println(len(raw), cap(raw), &raw[0]) //10000 10000 0xc00012c000
	return raw[:3]
}

//可以通过拷贝临时 slice 的数据，而不是重新切片来解决
func copySlice() []int {
	raw := make([]int, 10000)
	fmt.Printf("%T\n", raw)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]int, 3)
	copy(res, raw[:3])
	return res
}

func main() {
	data := get()
	fmt.Printf("%T\n", data)
	fmt.Println(len(data), cap(data), &data[0]) //3 10000 0xc00012c000


	fmt.Println("------------------")

	data1 := copySlice()
	fmt.Printf("%T\n", data1)
	fmt.Println(len(data1), cap(data1), &data1[0]) //3 10000 0xc00012c000

}
