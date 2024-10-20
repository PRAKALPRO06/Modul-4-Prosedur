package main

import "fmt"

func main() {
	var a, b, c, d int
	var permAC, combAC, permBD, combBD int

	fmt.Print("a: ")
	fmt.Scanln(&a)
	fmt.Print("b: ")
	fmt.Scanln(&b)
	fmt.Print("c: ")
	fmt.Scanln(&c)
	fmt.Print("d: ")
	fmt.Scanln(&d)

	if a >= c && b >= d {
		permutation(a, c, &permAC)
		combination(a, c, &combAC)
		fmt.Printf("%d, %d\n", permAC, combAC)

		permutation(b, d, &permBD)
		combination(b, d, &combBD)
		fmt.Printf("%d, %d\n", permBD, combBD)
	} else {
		fmt.Println("Invalid")
	}
}

func factorial(n int) int {
	var hasil int = 1
	var i int

	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}

	return hasil
}

func permutation(n, r int, result *int) {
	*result = factorial(n) / factorial(n-r)
}

func combination(n, r int, result *int) {
	*result = factorial(n) / (factorial(r) * factorial(n-r))
}
