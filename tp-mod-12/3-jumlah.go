package main

import "fmt"

func main() {
	var char byte
	var count int

	count = 0

	for char != '#' {
		count = count + 1
		fmt.Scanf("%c", &char)
	}

	fmt.Println(count - 1)
}
