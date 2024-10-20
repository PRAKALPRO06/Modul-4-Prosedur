package main

import "fmt"

func f1(x, y int) float64 { // pass by velue
	var hasil float64
	hasil = float64(2*x) - 0.5*float64(y) + 3.0
	return hasil
}

func f2(x, y int, hasil *float64) { //pas by reference
	*hasil = float64(2*x) - 0.5*float64(y) + 3.0
}

func main() {
	var a, b int
	var c float64

	// Take input for a and b
	fmt.Print("Enter two integers: ")
	fmt.Scan(&a, &b)

	// Call f2 to calculate and store result in
	f2(a, b, &c)

	// Print the result from f2
	fmt.Println("Result from f2 (stored in c):", c)

	// Call f1 and print the result
	resultF1 := f1(b, a)
	fmt.Println("Result from f1:", resultF1)
}
