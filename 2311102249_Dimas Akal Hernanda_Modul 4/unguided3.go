package main

import "fmt"

func deret(n int) {
	fmt.Print(n)
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(" ", n)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan positif: ")
	fmt.Scan(&n)

	// Memastikan bilangan yang dimasukkan adalah positif
	if n > 0 {
		deret(n)
	} else {
		fmt.Println("Silakan masukkan bilangan positif.")
	}
}
