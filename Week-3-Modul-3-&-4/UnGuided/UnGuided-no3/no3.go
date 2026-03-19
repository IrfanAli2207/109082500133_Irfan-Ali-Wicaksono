package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    for n != 1 {
        fmt.Print(n, " ")  
        
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
    }

    fmt.Print(1)
}