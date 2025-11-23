package main

import "fmt"

func main() {
	var a int
	var mon string

	fmt.Scan(&a)

	switch a {
	case 1:
		mon = "Januari"
	case 2:
		mon = "Februari"
	case 3:
		mon = "Maret"
	case 4:
		mon = "April"
	case 5:
		mon = "Mei"
	case 6:
		mon = "Juni"
	case 7:
		mon = "Juli"
	case 8:
		mon = "Agustus"
	case 9:
		mon = "September"
	case 10:
		mon = "Oktober"
	case 11:
		mon = "November"
	case 12:
		mon = "Desember"
	default:
		mon = ""
	}

	if mon != "" {
		fmt.Println(mon)
	} else {
		fmt.Println("Input is invalid")
	}
}
