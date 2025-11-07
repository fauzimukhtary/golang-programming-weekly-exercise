package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("\nINPUT :\n")
	fmt.Scan(&x, &y)
	function(x, y)
}

func function(x, y int) {
	a := float64(x)
	b := float64(y)
	f := ((22 / b) / (7 + (2 * a))) + (a * b)
	fmt.Printf("\nOUTPUT:\nMuhammad Fauzi Ramadhan | %g\n", f)
}
