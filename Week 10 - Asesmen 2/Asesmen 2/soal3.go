package main

import "fmt"

const n = 10

func main() {

	var provinsi [n]string
	var populasi [n]int
	var tumbuh [n]float64

	
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&provinsi[i], &populasi[i], &tumbuh[i])
	}


	max := tumbuh[0]
	indexMax := 0 

	for i := 1; i < n; i++ {
		if tumbuh[i] > max {
			max = tumbuh[i]
			indexMax = i
		}
	}

	
	var cari string

	fmt.Println()
	fmt.Print("Masukkan nama provinsi : ")
	fmt.Scan(&cari)


	indexCari := -1

	for i := 0; i < n; i++ {
		if provinsi[i] == cari {
			indexCari = i
		}
	}


	fmt.Println()
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", provinsi[indexMax])

	if indexCari != -1 {
		fmt.Println("Data provinsi yang dicari :", provinsi[indexCari])
	} else {
		fmt.Println("Provinsi tidak ditemukan")
	}

	fmt.Println()
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")

	for i := 0; i < n; i++ {

		if tumbuh[i] > 0.02 {

			prediksi := int((tumbuh[i] + 1) * float64(populasi[i]))

			fmt.Println(provinsi[i], prediksi)
		}
	}
}