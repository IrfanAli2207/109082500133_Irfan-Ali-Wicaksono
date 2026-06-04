package main

import "fmt"

func main() {
	var data []int
	var angka int

	for {
		fmt.Scan(&angka)

		if angka == -5313541 {
			break
		}

		if angka == 0 {
			n := len(data)

			for i := 0; i < n-1; i++ {
				min := i
				for j := i + 1; j < n; j++ {
					if data[j] < data[min] {
						min = j
					}
				}
				data[i], data[min] = data[min], data[i]
			}

			if n%2 == 1 {
				fmt.Println(data[n/2])
			} else {
				median := float64(data[n/2-1]+data[n/2]) / 2
				fmt.Println(median)
			}

		} else {
			data = append(data, angka)
		}
	}
}
