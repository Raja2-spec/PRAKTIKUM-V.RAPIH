package main

import (
        "fmt"
        "math"
)

func jarakDuaTitik(x1, y1, x2, y2 float64) float64 {
        // Menghitung jarak antara dua titik menggunakan teorema Pythagoras
        dx := x2 - x1
        dy := y2 - y1
        return math.Sqrt(dx*dx + dy*dy)
}

func main() {
        var xA, yA, xB, yB, xC, yC float64

        // Meminta input koordinat dari pengguna
        fmt.Println("Masukkan koordinat titik A (x y):")
        fmt.Scan(&xA, &yA)

        fmt.Println("Masukkan koordinat titik B (x y):")
        fmt.Scan(&xB, &yB)

        fmt.Println("Masukkan koordinat titik C (x y):")
        fmt.Scan(&xC, &yC)

        // Menghitung panjang sisi-sisi segitiga
        panjangAB := jarakDuaTitik(xA, yA, xB, yB)
        panjangBC := jarakDuaTitik(xB, yB, xC, yC)
        panjangCA := jarakDuaTitik(xC, yC, xA, yA)

        // Menentukan sisi terpanjang
        var sisiTerpanjang float64
        if panjangAB >= panjangBC && panjangAB >= panjangCA {
                sisiTerpanjang = panjangAB
        } else if panjangBC >= panjangAB && panjangBC >= panjangCA {
                sisiTerpanjang = panjangBC
        } else {
                sisiTerpanjang = panjangCA
        }

        // Menampilkan hasil
        fmt.Printf("Panjang sisi terpanjang segitiga adalah: %.2f satuan\n", sisiTerpanjang)
}
