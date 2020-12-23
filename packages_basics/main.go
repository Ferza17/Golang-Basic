package main

import (
	"fmt"
	"mypackages/numbers"
)

func main() {
	fmt.Printf("%d is even: %t \n", 40, numbers.Even(40))
	fmt.Println(numbers.IsPrime(13), numbers.IsPrime(100))

}
