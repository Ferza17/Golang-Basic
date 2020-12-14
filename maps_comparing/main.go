package main

import "fmt"

func main() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "X"}

	// fmt.Sprintf(a == b)

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps Are Equal!")
	} else {
		fmt.Println("Maps Aren't Equal!")
	}

}
