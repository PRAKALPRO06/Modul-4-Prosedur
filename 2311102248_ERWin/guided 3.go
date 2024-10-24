package main

import "fmt"

func f1(x, y int) float64 {
	var hasil float64
	hasil = float64(2*x) - 0.5*float64(y) + 3.0
	return hasil
}

func f2(x, y int, hasil *float64) { //pass by preference
	*hasil = float64(2*x) - 0.5*float64(y) + 3.0
}

func main() {
	var x, y int
	var z float64

	fmt.Print("Enter two intergers : ")
	fmt.Scan(&x, &y)

	f2(x, y, &z)
	fmt.Println("Result from f2 (stored in z) : ", z)

	resultF1 := f1(y, x)
	fmt.Println("Result from f1 : ", resultF1)
}
