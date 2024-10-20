package main

import (
	"fmt"
)

func faktorial(n int, result *int) {
	if n < 0 {
		*result = 0
		return
	}
	*result = 1

	for i := 1; i <= n; i++ {
		*result *= i
	}
}

func permutasi(n, r int, result *int) {
	if n < r {
		*result = 0
		return
	}
	var nFact, nrFact int
	faktorial(n, &nFact)
	faktorial(n-r, &nrFact)
	*result = nFact / nrFact
}

func kombinasi(n, r int, result *int) {
	if n < r {
		*result = 0
		return
	}
	var nFact, rFact, nrFact int
	faktorial(n, &nFact)
	faktorial(r, &rFact)
	faktorial(n-r, &nrFact)
	*result = nFact / (rFact * nrFact)
}

func main() {
	var a, b, c, d int
	var permA, combA, permB, combB int

	fmt.Print("Masukkan Input : ")
	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		fmt.Print("Input tidak valid, silakan masukkan input.")
		return
	}

	if a >= c && b >= d {

		permutasi(a, c, &permA)
		kombinasi(a, c, &combA)
		fmt.Printf("%d, %d\n", permA, combA)

		permutasi(b, d, &permB)
		kombinasi(b, d, &combB)
		fmt.Printf("%d, %d\n", permB, combB)
	} else {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d")
	}
}
