package main

import "fmt"

func main() {
	var p, l int

	fmt.Print("masukkan panjang (p) : ")
	fmt.Scan(&p)
	fmt.Print("masukkan lebar (l) : ")
	fmt.Scan(&l)	

	luas := p * l
	keliling := (2 * p) + (2 * l)

	fmt.Printf("luas persegi panjang : %d\n", luas)
	fmt.Printf("keliling persegi panjang : %d\n", keliling)

}
