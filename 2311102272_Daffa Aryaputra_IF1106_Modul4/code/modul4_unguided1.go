package main

import (
	"fmt"
)

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var factN, factNR int
	factorial(n, &factN)
	factorial(n-r, &factNR)
	*hasil = factN / factNR
}

func combination(n, r int, hasil *int) {
	var perm, factR int
	permutation(n, r, &perm)
	factorial(r, &factR)
	*hasil = perm / factR
}

func main() {
	var a, b, c, d int
	fmt.Println("Masukkan empat bilangan bulat positif (a b c d):")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		var permAC, combAC, permBD, combBD int
		permutation(a, c, &permAC)
		combination(a, c, &combAC)
		permutation(b, d, &permBD)
		combination(b, d, &combBD)

		fmt.Printf("%d %d\n", permAC, combAC)
		fmt.Printf("%d %d\n", permBD, combBD)
	} else {
		fmt.Println("Input tidak memenuhi syarat a >= c dan b >= d")
	}
}