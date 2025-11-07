package main

import "fmt"

func main() {
	var n int
	var a1, a2, a3 int
	var res1, res2 int

	fmt.Scan(&n)

	a1 = n / 100
	a2 = n / 10 % 10
	a3 = n % 10
	res1 = (a1 * 110000) + (a2 * 1100) + (a3 * 11)
	res2 = res1 * res1

	fmt.Println(res1, res2)
}
