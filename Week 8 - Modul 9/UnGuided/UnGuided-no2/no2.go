package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	var arr [100]int 

	
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan nilai ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSemua isi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	
	fmt.Println("Elemen indeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	
	fmt.Println("Elemen indeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Println("Elemen indeks kelipatan x:")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	
	var hapus int
	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&hapus)

	for i := hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	

	fmt.Println("Array setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	
	total := 0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	rata := float64(total) / float64(n)
	fmt.Println("Rata-rata:", rata)

	
	var jumlah float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		jumlah += selisih * selisih
	}

	varian := jumlah / float64(n)
	stdDev := math.Sqrt(varian)

	fmt.Println("Standar deviasi:", stdDev)
}