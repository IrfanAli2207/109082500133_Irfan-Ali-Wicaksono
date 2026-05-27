package main

import "fmt"

const max = 20

func main() {
	var suara [max + 1]int
	var input int
	var totalMasuk, suaraSah int

	for {
		fmt.Scan(&input)

		
		if input == 0 {
			break
		}

		totalMasuk++

	
		if input >= 1 && input <= 20 {
			suara[input]++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)

	
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d : %d\n", i, suara[i])
		}
	}
}