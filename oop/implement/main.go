package main

import "fmt"

/*
	go中的interface类型，是一组方法签名，当某一个类型为接口中的所有方法提供定义时，它被称为实现接口
	接口中指定了类型应该具有的方法，类型决定了如何实现这些方法

	//定义接口
	type interface_name interface {
		method_name1 [return_type]
		method_name2 [return_type]
		method_name3 [return_type]
		...
		method_namen [return_type]
	}

	//定义结构体
	type struct_name struct {
		//variables
	}

	实现接口方法
	func (struct_name_variable struct_name) method_name1() [return_type] {
		//方法实现
	}
	...
	func (struct_name_variable struct_name) method_namen() [return_type] {
		//方法实现
	}

*/

//定义接口
type Phone interface {
	call()
	sendMsg(msg string) bool //带参数带返回值
}

//定义结构体
type IPhone struct {
	name  string
	price float32
}

type Mi struct {
	name  string
	price float32
}

type Huawei struct {
	name  string
	price float32
}

//实现接口方法：语法见上
//实现接口：需要实现接口中的所有方法，go语言中的实现就是将方法与类型进行一个绑定
//IPhone实现无参Call方法
func (iphone IPhone) call() {
	fmt.Printf("用%s打电话\n", iphone.name)
}

//IPhone实现有参sendMsg方法
func (iphone IPhone) sendMsg(msg string) bool {
	fmt.Printf("%s发短信, 内容为: %s\n", iphone.name, msg)
	return true
}

//Mi实现无参Call方法
func (mi Mi) call() {
	fmt.Printf("用%s打电话\n", mi.name)
}

//M实现有参sendMsg方法
func (mi Mi) sendMsg(msg string) bool {
	fmt.Printf("%s发短信, 内容为: %s\n", mi.name, msg)
	return true
}

//Huawei实现无参Call方法
func (huawei Huawei) call() {
	fmt.Printf("用%s打电话\n", huawei.name)
}

//空接口定义
type Object interface {
}

func main() {
	// 由于IPhone实现了Phone接口的所有方法，因此IPone是对象可以直接赋值给Phone类型
	var iphone Phone = IPhone{"苹果", 5000}
	var mi Phone = Mi{"小米", 3000}
	iphone.call()
	iphone.sendMsg("jane,明天见 来自iphone")
	mi.call()
	mi.sendMsg("jane,明天见 来自iphone")

	//Cannot use 'Huawei{"华为", 3000}' (type Huawei) as type
	//Phone Type does not implement 'Phone' as some methods are missing: sendMsg(msg string) bool
	//华为没有实现sendMsg，不可以赋值给Phone
	//var huawei Phone = Huawei{"华为", 3000}

	// 空接口：没有任何方法的接口就是空接口，因此任何类型都实现了空接口，这个类似于Java中的Object类型，空接口可以存储任何类型的值
	var obj Object = IPhone{name: "iphone14", price: 8000}
	//var name1 = IPhone(obj).name
	//空接口中取值需要使用类型断言，这里的obj.(IPhone)就是告诉程序我们的obj是一个IPhone类型
	var name = obj.(IPhone).name
	fmt.Println("name:", name)

	// 空接口可以不用命名，直接写为interface{}, 因为空接口是一种数据类型
	var any interface{} = "123"
	// 空接口可以存储任何的数据类型，但是在做赋值的时候，需要类型断言转换
	var str string = any.(string)
	fmt.Printf("any类型：%T,any值：%s \n", any, any)
	fmt.Println(str)
}
