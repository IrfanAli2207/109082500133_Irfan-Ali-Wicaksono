package main

import "fmt"

func main() {
	var mahasiswa [3]string 

	var i int
	for i = 0; i < 3; i++ {
		fmt.Printf("Masukkan data index ke-%d : ", i)
		fmt.Scan(&mahasiswa[i])
	}

	
	var j int
	for j = 0; j < 3; j++ {
		fmt.Println("data ke-", j, " : ", mahasiswa[j])
	}

	
	fmt.Println("index [:3] : ", mahasiswa[:3]) 
	fmt.Println("index [:1] : ", mahasiswa[1:]) 

	fmt.Println(len(mahasiswa)) 

	
	mahasiswa[1] = "Regita"
	fmt.Println(mahasiswa[1])

	
	var indexHapus = 0
	var idx int
	for idx = indexHapus; idx < len(mahasiswa)-1; idx++ {
		mahasiswa[idx] = mahasiswa[idx+1]
	}
	mahasiswa[len(mahasiswa)-1] = ""

	fmt.Println(mahasiswa[:3])


	var cariNama string = "Regita"
	var found bool = false
	var k int
	for k = 0; k < 3; k++ {
		if mahasiswa[k] == cariNama {
			fmt.Printf("data ditemukan pada index ke-%d", k)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("data tidak ditemukan")
	}
}