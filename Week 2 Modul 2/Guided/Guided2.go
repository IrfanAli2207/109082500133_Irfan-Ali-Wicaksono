package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	for angka < 10 { 
		fmt.Println("Angka sekarang adalah: ", angka)
		angka++
	}
}