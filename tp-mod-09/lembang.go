package main

import "fmt"

func main() {
	var mhs, bus int
	fmt.Scan(&mhs)
	bus = mhs / 45
	if mhs%45 != 0 {
		bus = bus + 1
	}
	fmt.Printf("Diperlukan %d bus untuk tamasya ke Lembang\n", bus)
}
