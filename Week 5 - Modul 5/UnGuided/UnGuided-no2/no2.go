package main

import "fmt"

func bintang(n int, i int) {
	if i > n {
		return
	}

	for j := 1; j <= i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	bintang(n, 1)
}