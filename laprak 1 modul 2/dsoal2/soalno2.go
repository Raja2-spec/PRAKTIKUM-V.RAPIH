package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Print("masukkan jari jari (r) : ")
	fmt.Scan(&r)

	luas := math.Pi * math.Pow(r, 2)
	keliling := 2 * math.Pi * r

	fmt.Printf("luas lingkaran : %.2f\n", luas)
	fmt.Printf("keliling lingkaran : %.2f\n", keliling)
}