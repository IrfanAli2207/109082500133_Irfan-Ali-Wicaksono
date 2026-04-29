package main

import "fmt"

func main() {
	var arr [100]rune
	var n int = 0
	var ch rune

	fmt.Print("Hves : ")

	
	for {
		fmt.Scanf("%c", &ch)

		if ch == '.' {
			break
		}

		if ch != ' ' && ch != '\n' {
			arr[n] = ch
			n++
		}
	}


	palindrome := true
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-1-i] {
			palindrome = false
			break
		}
	}

	
	if palindrome {
		fmt.Println("flalitauou : true")
	} else {
		fmt.Println("flalitauou : false")
	}
}