package main

import (
	"fmt"
)

func factorial(n int, result *int) {
	*result = 1
	if n == 0 {
		return
	}
	for i := 2; i <= n; i++ {
		*result *= i
	}
}

func permutasi(n, r int, result *int) {
	var faktorialN, faktorialNR int
	factorial(n, &faktorialN)
	factorial(n-r, &faktorialNR)
	*result = faktorialN / faktorialNR
}

func kombinasi(n, r int, result *int) {
	var faktorialN, faktorialR, faktorialNR int
	factorial(n, &faktorialN)
	factorial(r, &faktorialR)
	factorial(n-r, &faktorialNR)
	*result = faktorialN / (faktorialR * faktorialNR)
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan 4 input(dibedakan dengan spasi) : ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		var permutasiA, kombinasiA, permutasiB, kombinasiB int

		permutasi(a, c, &permutasiA)
		kombinasi(a, c, &kombinasiA)

		permutasi(b, d, &permutasiB)
		kombinasi(b, d, &kombinasiB)

		fmt.Println("Keluaran:")
		fmt.Println(permutasiA, kombinasiA)
		fmt.Println(permutasiB, kombinasiB)
	} else {
		fmt.Println("Syarat tidak terpenuhi")
	}
}
