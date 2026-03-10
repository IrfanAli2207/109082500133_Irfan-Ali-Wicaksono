package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("Tahun :", tahun)
		fmt.Println("Kabisat : True")
	} else {
		fmt.Println("Tahun :", tahun)
		fmt.Println("Kabisat : False")
	}
}