package main

import (
	"fmt"
)

func factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

func permComb(n, r int) (int, int) {
	if n < r {
		return 0, 0
	}
	perm := factorial(n) / factorial(n-r)
	comb := perm / factorial(r)
	return perm, comb
}

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	Pac, Cac := permComb(a, c)
	Pbd, Cbd := permComb(b, d)

	fmt.Println(Pac, Cac)
	fmt.Println(Pbd, Cbd)
}
