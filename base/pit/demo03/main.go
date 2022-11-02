package main

import "fmt"

/*
	string 类型的值是常量，不可更改
*/
func main() {
	/*x := "text"
	x[0] = 'T' //Cannot assign to x[0]*/

	x := "text"
	fmt.Println(&x) //0xc000010240
	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x)  //我ext
	fmt.Println(&x) //0xc000010240

}
