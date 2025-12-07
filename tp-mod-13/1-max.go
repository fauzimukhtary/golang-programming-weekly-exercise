package main

import "fmt"

func main() {
	var num, max int
	fmt.Scan(&num)
	max = num

	for num != 0 {
		if num > max {
			max = num
		}
		fmt.Scan(&num)
	}

	fmt.Println(max)
}
