package main

import "fmt"

func main() {
	//int age = 30; //c++ way
	var age = 30
	fmt.Println("age : ", age)

	name := "Fery"
	/*
		_   is blank identifier user-defined name of the program
		components used for the identification purpose.
	*/
	_ = name
	// fmt.Println(name)

	s := "Learning Golang !"

	fmt.Println(s)

	name = "Reza"

	name1 := "Aditya"
	_ = name1

	// Multiple Declare Variable
	car, cost := "Toyota", 50000
	fmt.Println(car, cost)

	car, year := "ISUZU", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	// Multiple Variable
	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 6

	j, i = i, j // Swapping variable
	_, _ = i, j

	fmt.Println(i, j)

	sum := 5 + 2.3

	fmt.Println(sum)

}
