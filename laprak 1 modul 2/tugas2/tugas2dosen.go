package main

import (
	"fmt"
	"math"
)

func main() {
	var r, kelilinglingkaran, luaslingkaran float64
	const phi = 3.14

	fmt.Println("Masukan jari-jari")
	fmt.Scan(&r)

	luaslingkaran = phi * math.Pow(r, 2)
	kelilinglingkaran = 2 * phi * r

	fmt.Println("Hasil dari luas lingkaran adalah", luaslingkaran)
	fmt.Println("Hasil dari keliling lingkaran adalah", kelilinglingkaran)
}
