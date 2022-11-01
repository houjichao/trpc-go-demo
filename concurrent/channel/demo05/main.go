package main

import "fmt"

func printer(ch1 chan int) {

	for {
		data := <-ch1

		if data == 0 {
			break
		}
		fmt.Println(data)
	}

	ch1 <- 0
}

func main() {

	c := make(chan int)

	go printer(c)

	for i := 1; i < 10; i++ {
		c <- i
	}

	c <- 0

	<-c
}
