package main

import "fmt"

/*
	range中对map进行操作
*/
func main() {

	d := map[string]string{
		"jack":  "111",
		"pony":  "222",
		"eason": "333",
	}
	fmt.Println(d)
	fmt.Println("--------------------")

	//删除
	for k := range d {
		if k == "pony" {
			delete(d, k)
		}
	}
	fmt.Println(d)
	fmt.Println("--------------------")

	//新增
	for k, v := range d {
		d[v] = k
	}
	fmt.Println(d)
	fmt.Println("--------------------")

}
