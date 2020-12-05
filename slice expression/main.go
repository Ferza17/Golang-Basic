package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	// a[start:stop]

	b := a[1:3]
	fmt.Printf("%v, %T\n", b, b) // it will ber printed [2,3], []int

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2) // it wil be printed [2 3]

	fmt.Println(s1[2:]) // same as s1[2:leng(s1)]
	fmt.Println(s1[:3]) // same as s1[0:3]
	fmt.Println(s1[:])  // same as s1[0:len(s1)]

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)
}
