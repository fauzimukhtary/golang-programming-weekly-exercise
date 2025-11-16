package main

import "fmt"

func main() {
	var usn, pw string

	fmt.Scan(&usn, &pw)

	if usn == "admin" && pw == "12345" {
		fmt.Println("Login berhasil")
	} else {
		if usn != "admin" {
			fmt.Println("Username tidak ditemukan")
		}

		if pw != "12345" {
			fmt.Println("Password salah")
		}
	}
}
