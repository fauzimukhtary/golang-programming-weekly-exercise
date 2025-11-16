package main

import "fmt"

func main() {
	var member string
	var a, b, c, d, e int
	var ganjilAll, genapAll bool
	var discount, cashback float64

	fmt.Scan(&member, &a, &b, &c, &d, &e)

	discount = 0
	cashback = 0
	ganjilAll = (a%2)+(b%2)+(c%2)+(d%2)+(e%2) == 5
	genapAll = (a%2)+(b%2)+(c%2)+(d%2)+(e%2) == 0

	if ganjilAll {
		discount = discount + 1.7*float64(c+d+e)
	} else if genapAll {
		cashback = cashback + 3.1*float64(a+b+c)
	} else {
		discount = discount + 1.7*float64(c+d+e)
		cashback = cashback + 3.1*float64(a+b+c)
	}

	if member == "yes" {
		discount = discount + 0.15*discount
		cashback = cashback + 0.15*cashback
	}

	if discount > 35 {
		discount = 35
	}

	if cashback > 35 {
		cashback = 35
	}

	fmt.Printf("cashback %.2f, dan discount %.2f\n", cashback, discount)
}
