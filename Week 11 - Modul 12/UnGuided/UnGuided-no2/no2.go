package main

import "fmt"

const max = 20

func main() {
	var suara [max + 1]int
	var input int
	var totalMasuk, suaraSah int

	for {
		fmt.Scan(&input)

		// berhenti jika input 0
		if input == 0 {
			break
		}

		totalMasuk++

		// validasi suara
		if input >= 1 && input <= 20 {
			suara[input]++
			suaraSah++
		}
	}

	var ketua, wakil int
	var max1, max2 int

	for i := 1; i <= 20; i++ {

		// mencari suara terbanyak pertama
		if suara[i] > max1 {
			max2 = max1
			wakil = ketua

			max1 = suara[i]
			ketua = i

		} else if suara[i] > max2 && i != ketua {

			// mencari suara terbanyak kedua
			max2 = suara[i]
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}