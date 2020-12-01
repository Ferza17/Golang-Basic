package main

import "fmt"

func main() {
	var x = 3   // Int
	var y = 3.1 // float

	x *= int(y) // Convert y from float to int
	fmt.Println(x)
}
