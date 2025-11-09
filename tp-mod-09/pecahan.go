package main

import "fmt"

func main() {
	var alice, bob int
	var uangAlice, uangBob int

	fmt.Scan(&alice, &bob)
	uangAlice = alice / 2000
	uangBob = bob / 2000
	if alice%2000 != 0 {
		uangAlice = uangAlice + 1
	}
	if bob%2000 != 0 {
		uangBob = uangBob + 1
	}

	fmt.Printf("PT. Alice memerlukan %d lembar USD 2,000\n", uangAlice)
	
	fmt.Printf("PT. Bob memerlukan %d lembar USD 2,000\n", uangBob)
}
