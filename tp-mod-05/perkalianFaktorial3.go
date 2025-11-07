package main

import "fmt"

func main() {
	var hasil, n int

	fmt.Scan(&n)

	for hasil = 1; n > 0; n-- {
		hasil *= n
	}

	fmt.Println(hasil)
}
