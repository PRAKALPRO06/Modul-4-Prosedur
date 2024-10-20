package main

import (
	"fmt"
)

func cetakDeret(n int) {
	if n <= 0 {
		fmt.Println("Masukan harus berupa bilangan integer positif yang lebih besar dari 0.")
		return
	}

	fmt.Printf("%d ", n)
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal (n): ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error: Masukan tidak valid.")
		return
	}

	if n > 1000000 {
		fmt.Println("Error: Nilai maksimum yang diperbolehkan adalah 1000000.")
		return
	}

	cetakDeret(n)
}
