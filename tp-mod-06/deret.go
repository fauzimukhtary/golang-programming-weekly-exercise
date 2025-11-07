package main

import "fmt"

func main() {
	var a, b, n, res int
	fmt.Scan(&a, &b)
	for n = 1; n <= b; n++ {
		res += a * n
	}
	fmt.Println(res)
}
