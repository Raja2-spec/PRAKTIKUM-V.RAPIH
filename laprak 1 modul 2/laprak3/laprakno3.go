package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luaslingkaran float64
	const phi = 3.14
	fmt.Println("masukkan jari jari")
	fmt.Scan(&r)
	luaslingkaran = 4 * phi * math.Pow(r,2)

	fmt.Println("hasilnya adalah", luaslingkaran)
}