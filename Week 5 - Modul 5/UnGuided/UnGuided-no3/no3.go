package main

import "fmt"

func cariFaktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Println(i)
	}

	cariFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	cariFaktor(n, 1)
}