package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	c := make(chan int)

	// <- channerl operator

	// Send
	c <- 10

	// RECEIVE
	num := <-c
	fmt.Println(<-c)
	_ = num

	close(c)

}
