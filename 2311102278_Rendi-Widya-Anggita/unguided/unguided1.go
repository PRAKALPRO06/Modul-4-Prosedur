package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Println("Masukkan 4 bilangan : ")
	fmt.Scan(&a, &b, &c, &d)

	// Menghitung permutasi dan kombinasi
	pC := permutasi(a, c)
	pD := permutasi(b, d)
	kC := kombinasi(a, c)
	kD := kombinasi(b, d)

	//output
	fmt.Printf("Permutasi P(%d, %d): %d\n", a, c, pC)
	fmt.Printf("Kombinasi C(%d, %d): %d\n", a, c, kC)
	fmt.Printf("Permutasi P(%d, %d): %d\n", b, d, pD)
	fmt.Printf("Kombinasi C(%d, %d): %d\n", b, d, kD)
}
