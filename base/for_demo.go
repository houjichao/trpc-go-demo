package base

import "fmt"

/*
	for 循环学习
	for init; condition; post {}
	for condition {}
	for {}

	for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	for key, value := range oldMap {
    	newMap[key] = value
	}

*/
func ForDemo() {
	//求和
	sum := 10
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Println("key is: %d - value is: %f\n", key, value)
	}

	for key := range map1 {
		fmt.Println("key is %d\n", key)
	}

	for _, value := range map1 {
		fmt.Println("value is %f", value)
	}

}
