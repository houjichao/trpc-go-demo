package main

import "fmt"

/*
	Go语言提供了两种精度的浮点数 float32 和 float64
	这些浮点数类型的取值范围可以从很微小到很巨大。浮点数取值范围的极限值可以在 math 包中找到：
	常量 math.MaxFloat32 表示 float32 能取到的最大数值，大约是 3.4e38；
	常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308；
	float32 和 float64 能表示的最小值分别为 1.4e-45 和 4.9e-324。
	浮点数会有一个精度的限制。所以应该优先使用float64，
 */
func main() {
	var f float32 = 1677721666 // 1 << 24
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(f == f+1) // "true"!

	/*
		1.6777217e+09
		1.6777217e+09
		true
	*/

	fmt.Println("--------------------")

	var f1 float64 = 1677721666 // 1 << 24
	fmt.Println(f1)
	fmt.Println(f1 + 1)
	fmt.Println(f1 == f1+1) // "false"!

	/**
		1.677721666e+09
		1.677721667e+09
		false
	*/
}
