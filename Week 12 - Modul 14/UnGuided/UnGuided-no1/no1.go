package main

import "fmt"

const NMAX = 100000

var rumah [NMAX]int

func selectionSort(m int) {
	for i := 0; i < m-1; i++ {
		min := i

		for j := i + 1; j < m; j++ {
			if rumah[j] < rumah[min] {
				min = j
			}
		}

		rumah[i], rumah[min] = rumah[min], rumah[i]
	}
}

func main() {
	var n, m int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(m)

		for j := 0; j < m; j++ {
			fmt.Print(rumah[j], " ")
		}
		fmt.Println()
	}
}