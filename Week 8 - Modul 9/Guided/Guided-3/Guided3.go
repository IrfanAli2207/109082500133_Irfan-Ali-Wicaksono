package main

import "fmt"

func main() {
	var nilai []int 
	
	var jmlData int = 3
	var idx int
	for idx = 0; idx < jmlData; idx++ {
		var inputNilai int
		fmt.Printf("Masukkan nilai ke-%d : ", idx)
		fmt.Scan(&inputNilai)
		nilai = append(nilai, inputNilai)
	}
	fmt.Println()


	fmt.Println("Semua data:", nilai)
	fmt.Println("Slice [0:2]:", nilai[:2]) 
	fmt.Println("Panjang slice:", len(nilai))

	
	nilai[0] = 100
	fmt.Println("\nSetelah update:", nilai)

	
	var indexHapus int = 1
	nilai = append(nilai[:indexHapus], nilai[indexHapus+1:]...) 
	fmt.Println("\nSetelah delete:", nilai)
	fmt.Println()

	
	var cariNilai int = 75
	var found bool = false
	var i int

	for i = 0; i < len(nilai); i++ {
		if nilai[i] == cariNilai {
			fmt.Printf("Nilai %d ditemukan di index %d\n", cariNilai, i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Nilai tidak ditemukan")
	}
}
