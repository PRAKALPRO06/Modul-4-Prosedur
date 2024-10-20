package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r)
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	// Input
	var a, b, c, d int
	fmt.Println("Masukkan 4 bilangan (a b c d):")
	fmt.Scan(&a, &b, &c, &d)

	// Mengecek syarat a >= c dan b >= d
	if a >= c && b >= d {
		// Menghitung permutasi dan kombinasi a terhadap c
		perm_a_c := permutasi(a, c)
		comb_a_c := kombinasi(a, c)

		// Menghitung permutasi dan kombinasi b terhadap d
		perm_b_d := permutasi(b, d)
		comb_b_d := kombinasi(b, d)

		// Output
		fmt.Printf("Permutasi(a, c): %d, Kombinasi(a, c): %d\n", perm_a_c, comb_a_c)
		fmt.Printf("Permutasi(b, d): %d, Kombinasi(b, d): %d\n", perm_b_d, comb_b_d)
	} else {
		fmt.Println("Input tidak valid. Pastikan a >= c dan b >= d.")
	}
}
