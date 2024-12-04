package main

import "fmt"

func main() {
	var nama_tanaman string
	fmt.Scan(&nama_tanaman)

	switch nama_tanaman {
	case "nepenthes", "dresera":
		fmt.Println("termasuk tanaman karnivora")
	case "venus", "sarracenia":
		fmt.Println("termasuk tanaman karnivora")
		fmt.Println("tidak asli indonesia")
	default:
		fmt.Println("tidak termasuk tanaman karnivora")	
		
	}

}