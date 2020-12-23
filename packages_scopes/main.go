package main

import "fmt"

func main() {
	fmt.Println("Scope means name visibility")
	sayHello("Fery")
	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water boiling Point int Degress Fahrenheit: ", tf)
}
