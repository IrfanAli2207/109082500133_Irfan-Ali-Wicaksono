package main

import "fmt"

const NMAX = 719

type buku struct {
	id     int
	judul  string
	penulis string
	rating int
}

type DaftarBuku [NMAX]buku

func BuatBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].rating,
		)
	}
}

func CetakFavorit(pustaka DaftarBuku, n int) int {
	idx := 0

	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idx].rating {
			idx = i
		}
	}

	return idx
}

func TutupBuku(pustaka DaftarBuku, n int, x int) int {
	for i := 0; i < n; i++ {
		if pustaka[i].id == x {
			return i
		}
	}
	return -1
}

func CariBuku(pustaka DaftarBuku, n int, x string) int {
	for i := 0; i < n; i++ {
		if pustaka[i].judul == x {
			return pustaka[i].rating
		}
	}
	return -1
}

func main() {
	var pustaka DaftarBuku
	var n int
	var idCari int
	var judulCari string

	fmt.Scan(&n)

	BuatBuku(&pustaka, n)


	idx := CetakFavorit(pustaka, n)
	fmt.Println("Buku favorit:", pustaka[idx].judul)

	
	fmt.Scan(&judulCari)
	rating := CariBuku(pustaka, n, judulCari)
	fmt.Println("Rating:", rating)

	
	fmt.Scan(&idCari)
	idx = TutupBuku(pustaka, n, idCari)

	if idx != -1 {
		fmt.Println("Ditemukan:", pustaka[idx].judul)
	} else {
		fmt.Println("Buku tidak ditemukan")
	}
}