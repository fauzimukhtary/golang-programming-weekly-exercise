package main

import "fmt"

func main() {
	var usia int

	fmt.Scan(&usia)

	if usia >= 0 && usia <= 12 {
		fmt.Printf("Usia %d termasuk kategori anak-anak\n", usia)
	} else if usia >= 13 && usia <= 17 {
		fmt.Printf("Usia %d termasuk kategori remaja\n", usia)
	} else if usia >= 18 && usia <= 55 {
		fmt.Printf("Usia %d termasuk kategori dewasa\n", usia)
	} else if usia > 55 {
		fmt.Printf("Usia %d termasuk kategori lansia\n", usia)
	}
}
