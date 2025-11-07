package main

import "fmt"

func main() {
	var n, s, pu, su, sq, jj int
	var k, totalpu, totalsu, totalsq, totaljj float64

	fmt.Scan(&s)
	fmt.Scan(&pu, &su, &sq, &jj)

	for n = 1; n <= pu; n++ {
		totalpu += float64(n) * 0.5
	}

	for n = 1; n <= su; n++ {
		totalsu += float64(n) * 0.3
	}

	for n = 1; n <= sq; n++ {
		totalsq += float64(n) * 0.2
	}

	for n = 1; n <= jj; n++ {
		totaljj += float64(n) * 0.6
	}

	k = float64(s) * (totalpu + totalsu + totalsq + totaljj)
	fmt.Printf("Total kalori terbakar hari ini sebanyak %g kalori\n", k)
}
