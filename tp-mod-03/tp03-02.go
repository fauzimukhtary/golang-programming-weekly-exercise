package main

import "fmt"

func main() {
	var input, digit1, digit2, digit3 int
	
	fmt.Scanln(&input)
	
	digit1 = input / 100
	digit2 = input / 10 % 10
	digit3 = input % 10
	
	fmt.Println(digit1, digit2, digit3)
}
