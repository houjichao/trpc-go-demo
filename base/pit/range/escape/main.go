package main

import "fmt"

/*
	range中的坑
	变量逃逸
*/
type user struct {
	name string
	age  uint64
}

func main1() {
	u := []user{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	n := make([]*user, 0, len(u))
	for _, v := range u {
		n = append(n, &v)
	}
	fmt.Println(n)
	for _, v := range n {
		fmt.Println(v)
	}
}

func main() {
	u := []user{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	n := make([]*user, 0, len(u))
	fmt.Printf("%T \n", n)
	for _, v := range u {
		o := v
		n = append(n, &o)
		//n = append(n, &u[k])
	}
	fmt.Println(n)
	for _, v := range n {
		fmt.Println(*v)
	}
	/*
		在for range中，变量v是用来保存迭代切片所得的值，因为v只被声明了一次，每次迭代的值都是赋值给v，
		该变量的内存地址始终未变，这样讲他的地址追加到新的切片中，该切片保存的都是同一个地址，这肯定无法达到预期效果的。
		这里还需要注意一点，变量v的地址也并不是指向原来切片u[2]的，因我在使用range迭代的时候，变量v的数据是切片的拷贝数据，
		所以直接copy了结构体数据。
	*/

	//修改变量的问题
	for k, v := range u {
		if v.age != 18 {
			u[k].age = 18
		}
	}
	fmt.Println(u)

}
