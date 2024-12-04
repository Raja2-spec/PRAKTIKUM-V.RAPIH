package main

import "fmt"

func main() {
	var tipe_kendaraan string
	var durasi, tarif int

	fmt.Print("masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&tipe_kendaraan)
	fmt.Print("masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch {
	case tipe_kendaraan == "motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case tipe_kendaraan == "motor" && durasi > 2:
		tarif = 9000
	case tipe_kendaraan == "mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case tipe_kendaraan == "mobil" && durasi > 2:
		tarif = 20000
	case tipe_kendaraan == "truk" && durasi >=1 && durasi <= 2:
		tarif = 25000
	case tipe_kendaraan == "truk" && durasi >2:
		tarif = 35000
	default:
		fmt.Println("jenis kendaraan atau durasi tidak valid")
	
	}
	fmt.Printf("tarif parkir : Rp %d", tarif)
}