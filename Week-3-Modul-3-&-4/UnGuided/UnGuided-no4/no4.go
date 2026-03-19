package main

import (
	"fmt"
	"math"
)

func persegi(s int) {
	fmt.Println("Luas :", s*s)
	fmt.Println("Keliling :", 4*s)
}

func persegipanjang(p, l int) {
	fmt.Println("Luas :", p*l)
	fmt.Println("Keliling :", 2*(p+l))
}

func lingkaran(r float64) {
	fmt.Println("Luas :", math.Pi*r*r)
	fmt.Println("Keliling :", 2*math.Pi*r)
}

func main() {
	var pilih int

	fmt.Println("1. Persegi")
	fmt.Println("2. Persegi Panjang")
	fmt.Println("3. Lingkaran")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		var s int
		fmt.Print("Sisi: ")
		fmt.Scan(&s)
		persegi(s)

	case 2:
		var p, l int
		fmt.Print("Panjang: ")
		fmt.Scan(&p)
		fmt.Print("Lebar: ")
		fmt.Scan(&l)
		persegipanjang(p, l)

	case 3:
		var r float64
		fmt.Print("Jari-jari: ")
		fmt.Scan(&r)
		lingkaran(r)

	default:
		fmt.Println("Salah pilih")
	}
}