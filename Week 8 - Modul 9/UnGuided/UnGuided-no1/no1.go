	package main

	import "fmt"


	type Titik struct {
		x int
		y int
	}


	type Lingkaran struct {
		pusat  Titik
		radius int
	}

	func main() {
		
		var l1, l2 Lingkaran
		var t Titik

		
		fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)

		
		fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)

		
		fmt.Scan(&t.x, &t.y)

		

		
		jarakL1 := (t.x-l1.pusat.x)*(t.x-l1.pusat.x) + (t.y-l1.pusat.y)*(t.y-l1.pusat.y)
		diL1 := jarakL1 <= (l1.radius * l1.radius)

		
		jarakL2 := (t.x-l2.pusat.x)*(t.x-l2.pusat.x) + (t.y-l2.pusat.y)*(t.y-l2.pusat.y)
		diL2 := jarakL2 <= (l2.radius * l2.radius)

		
		
		if diL1 && diL2 {
			fmt.Println("Titik di dalam lingkaran 1 dan 2")
		} else if diL1 {
			fmt.Println("Titik di dalam lingkaran 1")
		} else if diL2 {
			fmt.Println("Titik di dalam lingkaran 2")
		} else {
			fmt.Println("Titik di luar lingkaran 1 dan 2")
		}
	}