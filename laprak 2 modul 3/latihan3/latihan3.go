package main

import "fmt"

type transaksi struct {
	nama_barang  string
	jumlah       int
	harga_satuan float64
	total_harga  float64
}

func main() {
	var t transaksi

	fmt.Println("informasi transaksi")

	fmt.Print("nama barang : ")
	fmt.Scanln(&t.nama_barang)

	fmt.Print("jumlah davi : ")
	fmt.Scanln(&t.jumlah)

	fmt.Print("harga davi : ")
	fmt.Scanln(&t.harga_satuan)

	t.total_harga = float64(t.jumlah) * float64(t.harga_satuan)

	fmt.Print("hasilnya adalah ", t.total_harga)

}
