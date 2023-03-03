package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type person struct {
	name string
	age  int
}

func (u *user) notify() {
	fmt.Println("sendNotify to user", u.name)
}

func (p *person) notify() {
	fmt.Println("sendNotify to person", p.name)
}

func main() {
	user1 := user{"张三", "qq@qq.com"}

	var n1 notifier
	n1 = &user1
	n1.notify()

	person1 := person{"李四", 20}

	var n2 notifier
	n2 = &person1
	n2.notify()
}
