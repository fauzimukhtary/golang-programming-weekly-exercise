package main

import "fmt"

func main() {
	var pound, kilogram, litre float64
	const poundToKilogram float64 = 0.45359237
	const litreToKilogram float64 = 0.80

	fmt.Scanln(&pound)

	kilogram = pound * poundToKilogram
	litre = kilogram / litreToKilogram

	fmt.Printf("%.6f kg %.6f L\n", kilogram, litre)
}
