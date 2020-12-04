package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i : %d\n", i)
	}

	// WHile loops doesnt exist in go then you can use this aproach
	j := 10
	for j >= 0 {
		fmt.Printf("j : %d\n", j)
		j--
	}
}
