package main

import (
	"fmt"
	"testing"
)

/*
在结构体中使用指针
*/
type receiver1 struct {
	Id   int
	Name string
	Age  int
}

type Student struct {
	No        int
	Map       map[string]int
	ReceiverA receiver1
	ReceiverB *receiver1
}

func (stu Student) updateA() {
	fmt.Printf(" %p\n", &stu)
	stu.Map["a"] = 1
	stu.ReceiverA = receiver1{Name: "ww"}
	stu.ReceiverB = &receiver1{Name: "ww"}
}

func (stu Student) updateB() {
	stu.Map["b"] = 2
	stu.ReceiverA = receiver1{Name: "ww"}
	stu.ReceiverB = &receiver1{Name: "ww"}
}

func TestPointerTest1(t *testing.T) {
	student := Student{Map: map[string]int{"S": 0}, ReceiverA: receiver1{Name: "A"}, ReceiverB: &receiver1{Name: "B"}}
	fmt.Printf(" %p\n", &student)
	//fmt.Println(student, *student.ReceiverB)//{0 map[S:0] {0 A 0} 0xc00005c060} {0 B 0}
	student.updateA()
	//fmt.Println(student, *student.ReceiverB)//{0 map[S:0 a:1] {0 A 0} 0xc00005c060} {0 B 0}
}
