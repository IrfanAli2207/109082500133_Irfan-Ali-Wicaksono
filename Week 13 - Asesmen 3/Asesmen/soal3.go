package main

import "fmt"

const NMAX = 100000

// struct partai
type partai struct {
	nama  int
	suara int
}

// array partai
type tabPartai [NMAX]partai

// sequential search
func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n int = 0
	var x int
	var idx int

	fmt.Println("Masukkan proses input suara:")

	// input sampai -1
	for {
		fmt.Scan(&x)

		if x == -1 {
			break
		}

		idx = posisi(p, n, x)

		if idx == -1 {
			// partai baru
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			// partai sudah ada
			p[idx].suara++
		}
	}

	// Insertion Sort Descending berdasarkan suara
	for i := 1; i < n; i++ {
		temp := p[i]
		j := i - 1

		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}

		p[j+1] = temp
	}

	// output
	fmt.Println("\nHasil Perhitungan suara :")

	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}

	fmt.Println()
}
