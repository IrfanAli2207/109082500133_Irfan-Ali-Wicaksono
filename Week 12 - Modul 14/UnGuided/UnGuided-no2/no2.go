package main

import "fmt"

const NMAX = 100000

func main() {
	var n, m, x int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var ganjil, genap [NMAX]int
		jmlGanjil := 0
		jmlGenap := 0

		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&x)

			if x%2 == 0 {
				genap[jmlGenap] = x
				jmlGenap++
			} else {
				ganjil[jmlGanjil] = x
				jmlGanjil++
			}
		}

		// Selection Sort ganjil naik
		for j := 0; j < jmlGanjil-1; j++ {
			min := j
			for k := j + 1; k < jmlGanjil; k++ {
				if ganjil[k] < ganjil[min] {
					min = k
				}
			}
			ganjil[j], ganjil[min] = ganjil[min], ganjil[j]
		}

		// Selection Sort genap turun
		for j := 0; j < jmlGenap-1; j++ {
			max := j
			for k := j + 1; k < jmlGenap; k++ {
				if genap[k] > genap[max] {
					max = k
				}
			}
			genap[j], genap[max] = genap[max], genap[j]
		}

		for j := 0; j < jmlGanjil; j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < jmlGenap; j++ {
			fmt.Print(genap[j], " ")
		}

		fmt.Println()
	}
}