package main

import "fmt"

func main() {
	var target, num int
	var result, counter int

	fmt.Scan(&target)
	fmt.Scan(&num)
	counter = 1
	result = 0

	for num != 0 {
		fmt.Scan(&num)
		counter = counter + 1
		if num == target {
			result = counter
		}
	}

	if result != 0 {
		fmt.Println(result)
	} else {
		fmt.Println("TIDAK ADA")
	}
}
