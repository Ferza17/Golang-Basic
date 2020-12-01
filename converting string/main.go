package main

import (
	"fmt"
	"strconv"
)

func main() {
	// s := string(99) // Printed c because is ASCI code for 99
	// fmt.Println(s)

	//  float
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	//  int
	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	// fmt.Println(string(34234))

	// String to Float
	s1 := "3.123"
	fmt.Printf("%T \n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err

	fmt.Println(f1 * 3.4)

	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, i value is %q\n", s2, s2)
}
