package main

import (
	"fmt"
	"reflect"
)

type data struct {
	num    int
	checks [10]func() bool   // 无法比较
	doIt   func() bool       // 无法比较
	m      map[string]string // 无法比较
	bytes  []byte            // 无法比较
}

type data1 struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

func main() {
	v1 := data{}
	v2 := data{}
	//Go 提供了一些库函数来比较那些无法使用 == 比较的变量，比如使用 "reflect" 包的 DeepEqual() ：
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2)) // true'

	v3 := data1{}
	v4 := data1{}
	fmt.Println("v1 == v2: ", v3 == v4) // true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1, m2)) // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// 注意两个 slice 相等，值和顺序必须一致
	fmt.Println("v1 == v2: ", reflect.DeepEqual(s1, s2)) // true

}
