package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	fmt.Print("masukkan x : ")
	fmt.Scan(&x)
	fmt.Print("masukkan y : ")
	fmt.Scan(&y)

	//nilai f(x, y)
	hasil := 1/(3*math.Pow(x, 2)) + 10*y + 7
	fmt.Printf("nilai f(%f, %f) adalah : %f\n", x, y, hasil)
	
}