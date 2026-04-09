package main

import "fmt"

const pi = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(kiri, kanan float64) {
	if kiri == kanan {
		fmt.Println("BALANCE")
	} else {
		fmt.Println("Selisih massa zat cair kiri dan kanan:", kiri-kanan)
	}
}

func main() {
	var r, tKiri, tKanan float64
	var mjKiri, mjKanan float64

	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)

	fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)

	kiri := massa(r, tKiri, mjKiri)
	kanan := massa(r, tKanan, mjKanan)

	display(kiri, kanan)
}