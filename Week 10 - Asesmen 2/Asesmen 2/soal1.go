package main

import "fmt"

func main() {
	var a [100]int
	var b [100]int
	var c [100]int
	var nA, nB, nC int
	var x int
	var ada bool

	for {
		fmt.Scan(&x)

		ada = false
		for i := 0; i < nA; i++ {
			if a[i] == x {
				ada = true
			}
		}

		if ada {
			break
		}

		a[nA] = x
		nA++
	}

	for {
		fmt.Scan(&x)

		ada = false
		for i := 0; i < nB; i++ {
			if b[i] == x {
				ada = true
			}
		}

		if ada {
			break
		}

		b[nB] = x
		nB++
	}

	for i := 0; i < nA; i++ {
		for j := 0; j < nB; j++ {
			if a[i] == b[j] {
				c[nC] = a[i]
				nC++
			}
		}
	}

	// output
	for i := 0; i < nC; i++ {
		fmt.Print(c[i], " ")
	}
}4
