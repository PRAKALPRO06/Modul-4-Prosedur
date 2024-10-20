package main

import (
        "fmt"
)

func factorial(n int) int {
        if n == 0 {
                return 1
        }
        return n * factorial(n-1)
}

func permutation(n, r int) int {
        return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
        return permutation(n, r) / factorial(r)
}

func main() {
        var a, b, c, d int

        fmt.Print("Masukkan nilai a, b, c, d: ")
        fmt.Scan(&a, &b, &c, &d)

        if a >= c && b >= d {
                p1 := permutation(a, c)
                c1 := combination(a, c)
                p2 := permutation(b, d)
                c2 := combination(b, d)

                fmt.Printf("P(%d,%d): %d\n", a, c, p1)
                fmt.Printf("C(%d,%d): %d\n", a, c, c1)
                fmt.Printf("P(%d,%d): %d\n", b, d, p2)
                fmt.Printf("C(%d,%d): %d\n", b, d, c2)
        } else {
                fmt.Println("Nilai a harus lebih besar sama dengan c, dan b harus lebih besar sama dengan d.")
        }
}