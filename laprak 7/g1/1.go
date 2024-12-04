package main

import "fmt"

func main() {
	var jam12,jam24 int
	var label string
	fmt.Scan(&jam24)
	switch {
	case jam24 == 0:
		jam12 = 12
		label = "am"
	case jam24 < 12:
		jam12 = jam24
		label = "am"
	case jam24 == 12:
		jam12 = 12
		label = "pm"
	case jam24 > 12:
		jam12 = jam24 - 12
		label = "pm"

	}

	fmt.Println(jam12, label)

}