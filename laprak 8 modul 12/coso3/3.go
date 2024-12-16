package main

import "fmt"

func main() {
	var A, s1, s2, B, temp int
	fmt.Scan(&A)
	s1 = 0
	s2 = 1
	B = 0
	for B < A {
		fmt.Print(s1, " ")
		temp = s1 + s2
		s1 = s2
		s2 = temp
		B = B + 1
	}
	
}