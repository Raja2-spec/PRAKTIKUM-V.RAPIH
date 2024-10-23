package main

import "fmt"

func main() {
	var p, l, L, K int
	fmt.Println("Masukan panjang")
	fmt.Scan(&p)

	fmt.Println("Masukan luas")
	fmt.Scan(&l)

	L = p * l
	fmt.Println("hasil dari luas", L)

	K = (2 * p) + (2 * l)
	fmt.Println("hasil dari keliling", K)
}
