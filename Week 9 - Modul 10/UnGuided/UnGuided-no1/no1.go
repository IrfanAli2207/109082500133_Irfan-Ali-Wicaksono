package main

import "fmt"

type arrBerat [1000]float64

func cariMinMax(arr arrBerat, n int) (float64, float64) {
	var min, max float64
	min = arr[0]
	max = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return min, max
}

func main() {
	var n int
	var berat arrBerat

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i)
		fmt.Scan(&berat[i])
	}

	min, max := cariMinMax(berat, n)

	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
