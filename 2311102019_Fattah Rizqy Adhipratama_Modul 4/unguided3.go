package main

import "fmt"

func cetakDeret(n int) {
    fmt.Printf("%d ", n)
    for n != 1 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
        fmt.Printf("%d ", n)
    }
    fmt.Println()
}

func main() {
    var n int
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&n)

    if n > 0 && n < 1000000 {
        cetakDeret(n)
    } else {
        fmt.Println("Masukkan angka antara 1 dan 1000000")
    }
}
