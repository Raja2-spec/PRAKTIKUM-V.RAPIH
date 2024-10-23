package main
import "fmt"

func main() {
    var r float64
	var c float64
	
	fmt.Println("masukan suhu celcius")
	fmt.Scan(&c)
	r = (4 / 5 * c)

	fmt.Println("hasil nya", r)
}
