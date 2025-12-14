package main

import "fmt"

func main() {
	var i, day, mon, yr int
	var batasHari, batasBulan int
	var hari string
	var kabisat bool

	fmt.Print("Passport 1: ")
	fmt.Scan(&hari)
	batasBulan, i = 12, 1

	for hari != "Exit" {
		fmt.Scan(&day, &mon, &yr)

		kabisat = (yr%4 == 0 && yr%100 != 0) || (yr%400 == 0)

		switch mon {
		case 1, 3, 5, 7, 8, 10, 12:
			batasHari = 31
		case 4, 6, 9, 11:
			batasHari = 30
		case 2:
			if kabisat {
				batasHari = 29
			} else {
				batasHari = 28
			}
		}

		switch hari {
		case "Jumat", "Kamis":
			day += 4
		default:
			day += 2
		}

		if day > batasHari {
			day -= batasHari
			mon++
		}

		if mon > batasBulan {
			mon -= 12
			yr++
		}

		fmt.Print("passport bisa diambil pada ")
		fmt.Printf("tanggal %d bulan %d tahun %d\n\n", day, mon, yr)

		i++
		fmt.Printf("Passport %d: ", i)
		fmt.Scan(&hari)
	}
}
