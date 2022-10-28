package main

import "fmt"

/*
	关键字： break
*/
func main() {
	//功能：求1-100的和，当和第一次超过300的时候，停止循环

	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Println(sum)
		if sum >= 300 {
			break
		}
	}
	fmt.Println("------ok")

}
