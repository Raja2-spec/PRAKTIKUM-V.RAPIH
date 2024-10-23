package main

import (
	"fmt"
	"math"
)

func main() {
	var fxy,x,y float64
	fmt.Scan(&x,&y)

	fxy = 1.00/(3.00*math.Pow(x,2)+10)+10*y+7

	fmt.Println("Hasilnya adalah", fxy)
}