package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	b = a
	for b > 1 {
		fmt.Print(b, " x ")
		b = b - 1
	}
	fmt.Println(1)
}