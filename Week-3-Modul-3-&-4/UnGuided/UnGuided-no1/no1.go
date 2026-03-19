package main

import "fmt"

func f(n int) int {
	r := 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}

func p(n, r int) int {
	return f(n) / f(n-r)
}

func c(n, r int) int {
	return f(n) / (f(r) * f(n-r))
}

func main() {
	var a, b, c1, d int
	fmt.Scan(&a, &b, &c1, &d)

	fmt.Println(p(a, c1), c(a, c1))
	fmt.Println(p(b, d), c(b, d))
}