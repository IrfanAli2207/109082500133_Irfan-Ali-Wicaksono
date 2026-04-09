package main
import "fmt"

func main() {
	var jumlah, lama int
	var tujuan string

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)


	if lama > 3 {
		lama = 3
	}


	biaya := 620000


	if tujuan == "mancanegara" {
		biaya = biaya * 3 / 2
	}

	total := jumlah * lama * biaya

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %d\n", total)
}