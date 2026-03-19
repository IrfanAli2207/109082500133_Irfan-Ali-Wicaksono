package main

import "fmt"

func biaya(jenis string, masuk, keluar int) int {
	jam := keluar - masuk

	if jenis == "motor" {
		if masuk < 17 {
			return jam * 4000
		}
		return jam * 5000
	} else if jenis == "mobil" {
		if masuk < 17 {
			return jam * 6000
		}
		return jam * 7000
	}
	return 0
}

func main() {
	var jenis string
	var masuk, keluar int
	total := 0
	no := 1

	for {
		fmt.Print("Kendaraan (motor/mobil/-): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Scan(&masuk, &keluar)

		b := biaya(jenis, masuk, keluar)
		fmt.Println("Biaya:", b)

		total += b
		no++
	}

	fmt.Println("Total:", total)
}