package main

import "fmt"

const nMAX = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrMahasiswa [nMAX]mahasiswa

// fungsi mencari nilai pertama
func nilaiPertama(data arrMahasiswa, n int, cari string) int {
	for i := 0; i < n; i++ {
		if data[i].NIM == cari {
			return data[i].nilai
		}
	}
	return -1
}


func nilaiTerbesar(data arrMahasiswa, n int, cari string) int {
	max := -1

	for i := 0; i < n; i++ {
		if data[i].NIM == cari {
			if data[i].nilai > max {
				max = data[i].nilai
			}
		}
	}

	return max
}

func main() {
	var data arrMahasiswa
	var n int
	var cari string

	
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&cari)

	
	pertama := nilaiPertama(data, n, cari)
	terbesar := nilaiTerbesar(data, n, cari)

	
	if pertama != -1 {
		fmt.Println("Nilai pertama dari NIM", cari, "adalah", pertama)
		fmt.Println("Nilai terbesar dari NIM", cari, "adalah", terbesar)
	} else {
		fmt.Println("NIM tidak ditemukan")
	}
}
