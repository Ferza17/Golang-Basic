package main

import "fmt"

func main() {
	s1 := "Hey There Go!"
	fmt.Println(s1)

	fmt.Println("He Says: \"Hello!\"")

	fmt.Println(`He Says: "Hello!"`)

	s2 := `I Like \n Go!` // raw string
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`
Price : 100
Brand : Nike
	`)

	fmt.Println(`C:\Users\Fery`)
	fmt.Println(`C:\\Users\\Fery`)

	var s3 = "I Love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0: ", s3[0]) // it printed 73 because I is 73 in ASCII

	// s3[5] = 'x' // Error cannot assign

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)

}
