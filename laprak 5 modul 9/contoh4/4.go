package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%7 == 0 {
		fmt.Println(n / 7)
	}else {
		fmt.Println((n / 7) + 1)

	}	
}