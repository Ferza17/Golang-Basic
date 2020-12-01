package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255

	x++ // overflow, x is 0
	fmt.Println(x)

	// var b int8 = 127
	// // Println call has possible formatting directive
	// fmt.Println("%d\n", b+1)

	f := float32(math.MaxFloat32)

	fmt.Println(f)

	f = f * 1.2 // f overflows to infinite !

	fmt.Println(f)

	// const i int8 = 300

	/*
		if ur aplication needs to work with very very big numbers that could
		overflow, you should use package called "math/big"
	*/
}
