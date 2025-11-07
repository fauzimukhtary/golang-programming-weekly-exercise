package main

import "fmt"

func main() {
	var i, n, angka1, angka2 int

	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Scan(&angka1, &angka2)
		fmt.Println(angka1 * angka2)
	}
}
