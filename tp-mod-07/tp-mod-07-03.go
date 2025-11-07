package main

import "fmt"

func main() {
	var a, b int
	var m1, m2, m3 bool
	var menang bool

	fmt.Scan(&a, &b)

	m1 = a == b
	m2 = a-b == 1 || b-a == 1
	m3 = a-b == 5 || b-a == 5
	menang = m1 || m2 || m3

	fmt.Println("Menang?", menang)
}
