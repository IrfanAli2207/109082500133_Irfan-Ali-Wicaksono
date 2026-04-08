package main

import "fmt"

func tampilkan(n int) {
	if n == 0 {
		return
	}

	tampilkan(n - 1)
	fmt.Println(n)
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	tampilkan(n)
}