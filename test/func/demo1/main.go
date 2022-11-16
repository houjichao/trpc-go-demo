package main

import "fmt"

func test(num int) {
	fmt.Println(num)
}

//既然函数是一种数据类型，可以当做形参传递调用
//定义一个函数，把另一个函数作为形参
func test02(num1 int, num2 float32, testFunc func(int)) {
	fmt.Println("----test02----")
}

//函数也可以起别名
type myFunc func(int)

func test03(num1 int, num2 float32, testFunc myFunc) {
	fmt.Println("----test03----")
}

//go支持对函数返回值命名，而且可以有多个返回值，这是java语言中没有的

//定义一个函数，求两个数的和、差
//传统写法：返回值和返回值类型对应，顺序不能乱
func calc01(num1 int, num2 int) (int, int) {
	result01 := num1 + num2
	result02 := num1 - num2
	return result01, result02
}

//升级写法：对函数返回值命名，里面顺序就无所谓了，顺序不用对应，注意是顺序，名字要一致
func calc02(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	res := 5
	return sum, res
}

func test04(num1 int, num2 int) (success map[string]interface{}, err error) {
	res := map[string]interface{}{
		"code":    "111",
		"message": "222",
	}
	fmt.Println(res)
	return
}

func main() {
	a := test
	//函数也是一种数据类型，可以赋值给变量
	fmt.Printf("a的数据类型是：%T，test的数据类型是%T \n", a, test)

	//通过该变量可以对函数直接进行调用
	a(10)

	var b = test
	b(20)

	//调用test02参数
	test02(10, 3.14, test)
	test02(10, 3.13, a)

	//go支持自定义数据类型
	//语法： type 自定义数据类型 数据类型
	//相当于给数据类型起别名
	//举个例子
	type myInt int
	var num1 myInt = 30
	fmt.Printf("num1的数据类型是:%T,值为:%v \n", num1, num1)
	var num2 int
	//num2 = num1这里不能直接赋值数据类型不一致
	num2 = int(num1)
	fmt.Println(num2)

	//调用test03
	test03(10, 3.12, a)

	v1, v2 := calc01(10, 3)
	fmt.Printf("v1对应的值为%v,v2对应的值为%v \n", v1, v2)

	v3, v4 := calc02(20, 5)
	fmt.Printf("v3对应的值为%v,v4对应的值为%v \n", v3, v4)

	v5, v6 := test04(1, 2)
	fmt.Printf("v5对应的值为%v,v6对应的值为%v \n", v5, v6)

}
