package main

import "fmt"

func main() {
	var usia int
	var kk bool
	fmt.Scan(&usia, &kk)

	if usia >= 17 && kk {
		fmt.Print("bisa membuat ktp")
		} else {
			fmt.Print("belum bisa membuat ktp")
		}
}