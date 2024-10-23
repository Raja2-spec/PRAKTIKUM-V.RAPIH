package main

import "fmt"

func main() {
	var F float64
	var c float64
	fmt.Println("masukan suhu celcius")
	fmt.Scan(&c)
	F = (c * 9 / 5) + 32

	fmt.Println("hasil nya", F)
}
