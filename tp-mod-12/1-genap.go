package main

import "fmt"

func main() {
	var x, counter int
	fmt.Scan(&x)
	counter = 0
	for x != 0 {
		counter = counter + ((x + 1) % 2)
		x = x / 10
	}
	fmt.Println(counter)
}
