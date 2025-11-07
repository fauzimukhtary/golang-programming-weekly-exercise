package main

import "fmt"

func main() {
	var (
		kar rune
		capital, lower bool
	)
	
	fmt.Scanf("%c", &kar)
	
	capital = kar >= 'A' && kar <= 'Z'
	lower = kar >= 'a' && kar <= 'z'
	
	fmt.Println(capital || lower)
}
