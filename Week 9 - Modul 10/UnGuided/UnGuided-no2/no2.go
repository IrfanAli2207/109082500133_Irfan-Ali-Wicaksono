package main

import "fmt"

type arrBerat [1000]float64

func main() {
	var x, y int
	var berat arrBerat

	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	var idx int = 0
	var rata float64 = 0

	for i := 0; i < jumlahWadah; i++ {
		var total float64 = 0

		for j := 0; j < y; j++ {
			total += berat[idx]
			idx++
		}

		fmt.Print(total, " ")
		rata += total
	}

	fmt.Println()

	rata = rata / float64(jumlahWadah)
	fmt.Println(rata)
}
