package main

import "fmt"

func main() {
	//对从动态语言转过来的开发者来说，简短声明很好用，这可能会让人误会 := 是一个赋值操作符。
	//如果你在新的代码块中像下边这样误用了 :=，编译不会报错，但是变量不会按你的预期工作：
	x := 1
	println(x) // 1
	println(&x)
	{
		println(x) // 1
		println(&x)
		x := 2
		println(&x)
		println(x) // 2    // 新的 x 变量的作用域只在代码块内部
	}
	println(x) // 1
	println(&x)

	var y interface{} = nil
	_ = y

	fmt.Println(x)

	//nil 是 interface、function、pointer、map、slice 和 channel 类型变量的默认初始值。
	//但声明时不指定类型，编译器也无法推断出变量的具体类型。
	/*var z = nil    // error: use of untyped nil
	_ = z*/

}
