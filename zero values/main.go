package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	// Convert Variable
	a = int(b)
	fmt.Println(a, b)

	// var x int
	// x = "5"

	// Zero Values because it doesnt have a value that assign into the variable
	var value int
	var price float64
	var name string
	var done bool

	fmt.Print(value, price, name, done)
}
