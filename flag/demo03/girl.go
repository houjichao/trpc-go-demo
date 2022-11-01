package main

//1.引库
import (
	"flag"
	"fmt"
)

/*
	通过flag制作命令行工具
*/

//2.定义变量，该变量的作用是存储命令行参数传来的值
var height int

func main() {
	//3.配置命令信息，在main函数或者init函数中
	//param1:定义的变量引用，param2：命令的名称 3：命令参数的默认值，4.命令的用法提示
	flag.IntVar(&height, "height", 140, "身高")
	//4.解析参数
	flag.Parse()
	//5.打印变量
	fmt.Println("恭喜你获得了身高", height, "的女朋友")

	//go build -o girl girl.go
	/*
		$ ./girl -height 170
		$ ./girl --height 170
		$ ./girl --height=170
		$ ./girl -height=170
	*/

}
