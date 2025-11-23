package main

import "fmt"

func main() {
	var a, b, op int

	fmt.Scan(&a, &b, &op)

	if b == 0 && op >= 4 {
		fmt.Println("tidak terdefinisi")
	} else {
		switch op {
		case 1:
			fmt.Println(a + b)
		case 2:
			fmt.Println(a - b)
		case 3:
			fmt.Println(a * b)
		case 4:
			fmt.Println(float64(a) / float64(b))
		case 5:
			fmt.Println(a % b)
		case 6:
			fmt.Println(a / b)
		}
	}
}
