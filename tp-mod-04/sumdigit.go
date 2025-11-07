package main

import "fmt"

func main() {
	var a, b, d1, d2, d3, d4 int
	
	fmt.Scan(&a)
	
	d1 = a / 1000
	d2 = a / 100 % 10
	d3 = a / 10 % 10
	d4 = a % 10
	b = d1 * 1000 + d4 * 100 + d3 * 10 + d2
	
	fmt.Println(b, a + b)
}
