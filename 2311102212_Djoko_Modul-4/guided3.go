package main

import (
	"fmt"
)

func f1(x, y int) float64 {
	return float64(2*x) - 0.5*float64(y) + 3.0
}

func f2(x, y int, hasil *float64) {
	*hasil = float64(2*x) - 0.5*float64(y) + 3.0
}

func main() {
	var a, b int
	var c float64

	fmt.Println("Enter two integers:")
	fmt.Scan(&a, &b)

	f2(a, b, &c)
	fmt.Printf("Result from f2 (stored in c): %.2f\n", c)

	resultF1 := f1(b, a)
	fmt.Printf("Result from f1: %.2f\n", resultF1)
}
