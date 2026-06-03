package main
import "fmt"

type mahasiswa struct {
	nama, nim  string
	IPK float64
}

type arrMhs [100]mahasiswa

func insertionSortStruct(T *arrMhs, n int) {
	var temp mahasiswa
	var i, j int

	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1
		for j >= 0 && T[j].nama > temp.nama {	
			T[j+1] = T[j]	
			j = j - 1
		}
		T[j+1] = temp
	}
}

func main() {
	var data arrMhs
	var n, i int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa indeks ke-%d:\n", i)

		fmt.Print("Nama: ")
		fmt.Scan(&data[i].nama)

		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)

		fmt.Print("IPK: ")
		fmt.Scan(&data[i].IPK)
	}
	fmt.Println("\nData mahasiswa sebelum disorting:")
	for i = 0; i < n; i++ {
		fmt.Printf("Nama: %s, NIM: %s, IPK: %.2f\n", data[i].nama, data[i].nim, data[i].IPK)
	}

	insertionSortStruct(&data, n)

	fmt.Println("\nData mahasiswa setelah disorting:")
	for i = 0; i < n; i++ {
		fmt.Printf("Nama: %s, NIM: %s, IPK: %.2f\n", data[i].nama, data[i].nim, data[i].IPK)
	}
}