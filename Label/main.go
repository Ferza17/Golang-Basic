package main

import "fmt"

func main() {
	outer := 19
	_ = outer
	people := [5]string{"Fery", "Reza", "Aditya", "Johny", "Uye"}
	friends := [2]string{"Johny", "Simalakama"}
	// Labels doesnt conflict with any names of variable example outer
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}

	fmt.Println("Next instruction after the break")
}
