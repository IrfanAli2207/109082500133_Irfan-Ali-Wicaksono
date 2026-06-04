package main

import "fmt"

type Pemain struct {
	nama   string
	gol    int
	assist int
}

func main() {
	var n int
	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	var pemain [101]Pemain

	for i := 0; i < n; i++ {
		fmt.Scan(&pemain[i].nama, &pemain[i].gol, &pemain[i].assist)
	}

	// Selection Sort
	for i := 0; i < n-1; i++ {
		max := i

		for j := i + 1; j < n; j++ {
			if pemain[j].gol > pemain[max].gol {
				max = j
			} else if pemain[j].gol == pemain[max].gol {
				if pemain[j].assist > pemain[max].assist {
					max = j
				}
			}
		}

		pemain[i], pemain[max] = pemain[max], pemain[i]
	}

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Println(pemain[i].nama, pemain[i].gol, pemain[i].assist)
	}
}
