package main

import "fmt"

func main() {
	var a, i int
	for i = 0; i < 3; i++ {
		fmt.Scan(&a)
		if a < 0 {
			a = -a
		}
		fmt.Printf("%d ", a)
	}
}
