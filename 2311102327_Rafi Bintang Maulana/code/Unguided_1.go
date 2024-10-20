package main

import (
	"fmt"
)

func RumusFaktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * RumusFaktorial(n-1)
}

func RumusPermutasi(n, r int) int {
	return RumusFaktorial(n) / RumusFaktorial(n-r)
}

func RumusKombinasi(n, r int) int {
	return RumusFaktorial(n) / (RumusFaktorial(r) * RumusFaktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan empat bilangan (a b c d): ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Printf("========================================================\n")
		pa := RumusPermutasi(a, c)
		ca := RumusPermutasi(a, c)
		fmt.Printf("Permutasi(a dan c): %d, Kombinasi(a dan c): %d\n", pa, ca)
		fmt.Printf("--------------------------------------------------------\n")
		pb := RumusPermutasi(b, d)
		cb := RumusKombinasi(b, d)
		fmt.Printf("Permutasi(b, d): %d, Kombinasi(b, d): %d\n", pb, cb)
		fmt.Printf("========================================================\n")
	} else {
		fmt.Println("Input Salah. Pastikan a >= c dan b >= d.")
	}
}