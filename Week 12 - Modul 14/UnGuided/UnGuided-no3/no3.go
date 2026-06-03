package main

import "fmt"

const NMAX = 100

func main() {
	var data [NMAX]int
	var n, x int

	fmt.Scan(&x)
	for x >= 0 {
		data[n] = x
		n++
		fmt.Scan(&x)
	}

	// Insertion Sort
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && data[j] > temp {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}

	// Cetak array yang sudah terurut
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	// Cek jarak
	jarak := data[1] - data[0]
	tetap := true

	for i := 2; i < n; i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
		}
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}