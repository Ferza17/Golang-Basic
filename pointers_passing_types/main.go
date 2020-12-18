package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // Equivalent to p.price = 300
	(*p).productName = "Bicycle"
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValues(): ", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling changeValues(): ", quantity, price, name, sold)

	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling changeValuesByPointer(): ", quantity, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Watch",
	}

	changeProduct(gift)
	fmt.Println(gift)
	fmt.Println("BEFORE calling changeProductByPointer(): ", gift)
	changeProductByPointer(&gift)
	fmt.Println("AFTER calling changeProductByPointer(): ", gift)

}
