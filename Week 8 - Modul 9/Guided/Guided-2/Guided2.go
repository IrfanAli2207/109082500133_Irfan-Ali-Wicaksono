package main

import "fmt"

func main() {
	
	var nilai map[string]int = make(map[string]int)

	
	nilai["Dhimas"] = 90 
	nilai["Hafizh"] = 88 
	nilai["Regita"] = 95 

	
	fmt.Println("Data nilai : ")
	var nama string                 
	var grade int                   
	for nama, grade = range nilai { 
		fmt.Println(nama, " = ", grade)
	}
	fmt.Println()

	
	fmt.Println("Setelah Update : ")
	nilai["Regita"] = 92 
	print("Regita = ", nilai["Regita"])
	fmt.Println()
	fmt.Println()

	
	delete(nilai, "Dhimas") 
	fmt.Println("Data nilai setelah delete : ")
	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}
	fmt.Println()

	
	fmt.Println("Hasil Searching : ")
	var cariNama string = "Hafizh" 
	var isiValue int               
	var ok bool                    
	isiValue, ok = nilai[cariNama] 
	if ok {                       
		fmt.Println("Nilai", cariNama, " = ", isiValue)
	} else { 
		fmt.Println("Data tidak ditemukan")
	}
}
