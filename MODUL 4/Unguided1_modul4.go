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
	var factN, factR, factNR int
	factorial(n, &factN)
	factorial(r, &factR)
	factorial(n-r, &factNR)
	*hasil = factN / (factR * factNR)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var permA, combA, permB, combB int

	permutation(a, c, &permA)
	combination(a, c, &combA)
	permutation(b, d, &permB)
	combination(b, d, &combB)

	fmt.Printf("%d %d\n", permA, combA)
	fmt.Printf("%d %d\n", permB, combB)
}
