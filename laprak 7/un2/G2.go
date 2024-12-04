package main

import "fmt"

func main() {
	var jenisKendaraan string
	fmt.Scanf("%s",&jenisKendaraan)
	var jam int
	fmt.Scan(&jam)
	 switch{
	 case jenisKendaraan == "motor" :
		hasil := jam * 2000
		fmt.Println(hasil)
	case jenisKendaraan == "mobil" :
		hasil := jam * 5000
		fmt.Println(hasil)
		default :
		hasil := jam * 8000
		fmt.Println(hasil)
		
	 }
}