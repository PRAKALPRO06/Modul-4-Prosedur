package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func Permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func combination(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan nilai a (bilangan asli): ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b (bilangan asli): ")
	fmt.Scan(&b)

	fmt.Print("Masukkan nilai c (dengan syarat a >= c): ")
	fmt.Scan(&c)

	fmt.Print("Masukkan nilai d (dengan syarat b >= d): ")
	fmt.Scan(&d)

	fmt.Println("Keluaran:")
	fmt.Println("Permutasi(a, c):", Permutasi(a, c), "Kombinasi(a, c):", combination(a, c))
	fmt.Println("Permutasi(b, d):", Permutasi(b, d), "Kombinasi(b, d):", combination(b, d))
}
