package main

import "fmt"

func main() {
	var i, j, x, n int
	var total float64

	fmt.Scan(&x, &n)

	for i, j = 1, n; i <= n; i++ {
		total = total + (float64(j) / float64(i*x))
		j = j - 1
	}

	fmt.Printf("%.3f\n", total)
}
