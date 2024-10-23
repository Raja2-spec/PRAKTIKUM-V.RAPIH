package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("masukkan berat badan (kg): ")
	fmt.Scan(&beratBadan)
	fmt.Print("masukkan tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("BMI anda: %.2f", bmi)

}
