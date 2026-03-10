package main

import "fmt"

func main() {
	var pilihan, angka int

	fmt.Println("-- MENU --")
	fmt.Println("1. Cek positif / negatif")
	fmt.Println("2. Cek genap / ganjil")

	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&angka)

		if angka > 0 {
			fmt.Println("Angka positif")
		} else if angka < 0 {
			fmt.Println("Angka negatif")
		} else {
			fmt.Println("Angka nol")
		}

	case 2:
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&angka)

		if angka%2 == 0 {
			fmt.Println("Angka genap")
		} else {
			fmt.Println("Angka ganjil")
		}

	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}