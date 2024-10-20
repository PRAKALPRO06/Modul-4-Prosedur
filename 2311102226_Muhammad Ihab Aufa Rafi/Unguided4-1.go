package main

import "fmt"

var a, b, c, d int

func faktorial(n int, hasil *int) {
    *hasil = 1
    for i := 1; i <= n; i++ {
        *hasil *= i
    }
}

func permutasi(n, r int, hasil *int) {
    var faktN, faktNR int
    faktorial(n, &faktN)
    faktorial(n-r, &faktNR)
    *hasil = faktN / faktNR
}

func kombinasi(n, r int, hasil *int) {
    var faktN, faktR, faktNR int
    faktorial(n, &faktN)
    faktorial(r, &faktR)
    faktorial(n-r, &faktNR)
    *hasil = faktN / (faktR * faktNR)
}

func main() {
    fmt.Printf("Masukkan Input (dipisahkan oleh spasi) = ")
    fmt.Scan(&a, &b, &c, &d)

    if a >= c && b >= d {
        var perm1_226, komb1_226, perm2_226, komb2_226 int
        
        permutasi(a, c, &perm1_226)
        kombinasi(a, c, &komb1_226)
        fmt.Printf("%d, %d\n", perm1_226, komb1_226)
        
        permutasi(b, d, &perm2_226)
        kombinasi(b, d, &komb2_226)
        fmt.Printf("%d, %d\n", perm2_226, komb2_226)
    } else {
        fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d")
    }
}
