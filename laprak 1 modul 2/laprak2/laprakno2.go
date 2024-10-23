package main

import "fmt"

func main() {
	var (
		nama  string
		prodi = "S1-IF"
		kelas string
		nim   int
	)

	fmt.Println("masukkan nama")
	fmt.Scan(&nama)

	fmt.Println("masukkan kelas")
	fmt.Scan(&kelas)

	fmt.Println("masukan nim")
	fmt.Scan(&nim)
	
	fmt.Println("Pernalkan nama saya ", nama, "salah satu mahasiswa prodi ", prodi , "dari kelas ", kelas, "dengan nim ", nim)
}