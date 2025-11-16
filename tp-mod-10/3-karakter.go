package main

import "fmt"

func main() {
	var char byte

	fmt.Scanf("%c", &char)

	if char >= 'A' && char <= 'Z' {
		fmt.Println("Huruf Besar")
	} else if char >= 'a' && char <= 'z' {
		fmt.Println("Huruf Kecil")
	} else if char >= '0' && char <= '9' {
		fmt.Println("Bilangan")
	} else {
		fmt.Println("Simbol")
	}
}
