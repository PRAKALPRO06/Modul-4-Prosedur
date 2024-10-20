package main

import "fmt"

func f1(x, y int) float64 { //pass by value
	var hasil float64
	hasil = float64(2*x) - 0.5*float64(y) + 3.0
	return hasil
}

func f2(x, y int, hasil *float64) { //pass by referenc
	*hasil = float64(2*x) - 0.5*float64(y) + 3.0 //pass by referenc
}

func main() {
	var a, b int
	var c float64

	//Take input for a and b
	fmt.Print("Enter two integers : ")
	fmt.Scan(&a, &b)

	f2(a, b, &c)
	fmt.Println("Result from f2 (stored in c):", c)

	resultF1 := f1(b, a)
	fmt.Println("Result from f1", resultF1)
}
