package main

import "fmt"

func main() {
	language := "Golang"

	switch language {
	case "Phyton":
		fmt.Println("You are learning Phyton")
	case "Go", "Golang":
		fmt.Println("You are Learning Golang")
	default:
		fmt.Println("Any other Programming languange is good start!")
	}
}

// You dont need break statement because go automatically run break in case
