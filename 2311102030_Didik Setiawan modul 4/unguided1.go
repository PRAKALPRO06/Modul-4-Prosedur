package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan 4 input(dibedakan dengan spasi) : ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {

		permutasiA := permutasi(a, c)
		kombinasiA := kombinasi(a, c)

		permutasiB := permutasi(b, d)
		kombinasiB := kombinasi(b, d)

		fmt.Println("Keluaran:")
		fmt.Println(permutasiA, kombinasiA)
		fmt.Println(permutasiB, kombinasiB)
	} else {
		fmt.Println("Syarat tidak terpenuhi")
	}
}
