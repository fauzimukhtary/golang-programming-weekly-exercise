package main

import "fmt"

func main() {
	var qty, res int
	var member string
	var disc float64

	fmt.Scan(&qty, &member)
	disc = 0

	if member == "A" {
		disc += 0.1
	} else if member == "B" {
		disc += 0.05
	}

	if qty > 10 {
		disc *= 3
	} else if qty >= 5 {
		disc *= 2
	}

	if member == "C" && qty > 10 {
		disc = (disc * 0) + 0.1
	} else if member == "C" && qty >= 5 {
		disc = (disc * 0) + 0.05
	} else if member == "N" && qty > 10 {
		disc = (disc * 0) + 0.05
	}

	res = qty * int(10000*(1.0-disc))

	fmt.Printf("Rp %d\n", res)
}
