package main

import (
	"fmt"

	// import "fmt"  // Error

	f "fmt"
)

// Permitted in go

const done = false // Package scoped

func main() {
	var task = "Running" // local (block scoped)
	fmt.Println(task, done)
	f.Println("Bye !")
}
