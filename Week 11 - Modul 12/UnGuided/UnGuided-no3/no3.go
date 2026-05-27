package main

import "fmt"

const NMAX = 100000

var data [NMAX]int

// prosedur untuk mengisi array
func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

// fungsi mencari posisi data
func posisi(n, k int) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func main() {
	var n, k int

	// input n dan k
	fmt.Scan(&n, &k)

	// input array
	isiArray(n)

	// mencari posisi data
	hasil := posisi(n, k)

	// output
	if hasil != -1 {
		fmt.Println(hasil)
	} else {
		fmt.Println("TIDAK ADA")
	}
}