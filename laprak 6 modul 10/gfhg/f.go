package main

import (
	"fmt"
)

func main() {
	var ipk, tak float64
	fmt.Print("Masukkan IPK: ")
	fmt.Scanln(&ipk)
	fmt.Print("Masukkan Nilai TAK (jika berlaku): ")
	fmt.Scanln(&tak)

	var predicate string
	if ipk < 2.00 {
		predicate = "Poor"
	} else if ipk >= 2.00 && ipk <= 2.75 {
		predicate = "Fair"
	} else if ipk > 2.75 && ipk <= 3.00 {
		predicate = "Satisfactory"
	} else if ipk > 3.00 && ipk <= 3.50 {
		predicate = "Very Good"
	} else {
		predicate = "Excellent"
	}

	// Output
	fmt.Println("Predikat TAK:", predicate)
}