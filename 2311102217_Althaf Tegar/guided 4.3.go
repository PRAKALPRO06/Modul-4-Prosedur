package main

import "fmt"

func f1(x, y int) float64 {
	var hasil float64
	hasil = float64(2*x) - 0.5*float64(y) + 3.0
	return hasil
}

func f2(x, y int, hasil *float64) { //pass by reference
	*hasil = float64(2*x) - 0.5*float64(y) + 3.0 //pass by value
}

func main() {
	var a, b int
	var c float64

	//take input for a and b
	fmt.Print("Enter two integers: ")
	fmt.Scan(&a, &b)

	//call f2 to calculate and store the result in c
	f2(a, b, &c)

	//Print the result from f2
	fmt.Println("Result from f2 (stored in c): ", c)

	//call f1 and print the result
	resultF1 := f1(b, a) //(6 4)
	fmt.Println("Result from f1:", resultF1)
}
