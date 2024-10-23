package main

import "fmt"
        
const PI float32 = 3.14

type tabung struct {
	tinggi int
	luas, volume float64
	luasAlas, luasDinding float64
}

func main() {
	var t tabung
	fmt.Scan(&t.jari2, &t.tinggi)
	t.luasAlas = PI * (float64(t.jari2) * float64(t.jari2))
	t.luasDinding = float32 (t.tinggi) * (2 * PI * float32(t.jari2))
	t.luas = 2*t.luasAlas + t.luasDinding
	t.volume = t.luasAlas * float64 (t.tinggi)

	fmt.Println(t.luas, t.volume)

}