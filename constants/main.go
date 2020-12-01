package main

import (
	"fmt"
)

func main() {
	const days int = 7

	// contains error if u using const u must assign value
	//const pi int

	const n1, n2 = 1, 2

	// if u not assign value like min2 and min3, a value ( -500 )  will be repeat
	const (
		min1 = -500
		min2
		min3
	)

	var i int
	fmt.Println(i)
	fmt.Println(min1, min2, min3)

	/*
		======== Constants Rules =============

		1) you can't change a constant
		example:
		const pi float64 = 3.14
		pi = 7777

		2) You can't initiate a constant at runtime
		example:
		const power = math.Pow(2,3)

		3) You can't use a variable to initialize a constant
		example:
		t:5
		const tc = t
	*/

}
