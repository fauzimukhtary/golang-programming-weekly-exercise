package main

import "fmt"

func main() {
	var i, speed int
	var counter, total, grandTotal int
	var berhenti bool
	var avg float64

	counter, grandTotal, berhenti = 0, 0, false

	for {
		for i = 0; i < 3; i++ {
			total = 0
			fmt.Scan(&speed)
			if speed >= 90 {
				berhenti = true
			}
			total += speed
			grandTotal += speed
			counter++
		}
		if berhenti || total >= 210 {
			break
		}
	}

	avg = float64(grandTotal) / float64(counter)

	fmt.Printf(
		"Total Set: %d (rata-rata: %.2f WPM)",
		(counter / 3),
		avg,
	)
}
