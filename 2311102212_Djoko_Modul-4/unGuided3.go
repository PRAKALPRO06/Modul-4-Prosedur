package main

import (
	"fmt"
)

func main() {
	var sukuAwal int

	fmt.Print("Masukkan nilai dari 0 - 1000000: ")
	fmt.Scan(&sukuAwal)
	cetakDeret(sukuAwal)
}

func cetakDeret(n int) {

	var genap, ganjil int

	fmt.Print(n)
	for {
		if n%2 == 0 {
			genap = n / 2 
			n = genap
			fmt.Print(" ", n)
		} else {
			ganjil = 3*n + 1
			n = ganjil
			fmt.Print(" ", n)
		}
		if n == 1 {
			break
		}
	}
	fmt.Println() 
}

