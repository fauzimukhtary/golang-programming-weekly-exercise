package main

import "fmt"

func main() {
	var a, n int

	fmt.Scan(&n)

	for a = 1; a <= n; n-- {
		fmt.Print(n)
		fmt.Print(" x ")
	}

	fmt.Println("\b\b", " ")
}
