package main

import (
	"fmt"
)

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println("1")
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan awal (1-999999): ")
	fmt.Scan(&n)
	
	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Masukan tidak valid. Masukkan bilangan antara 1 dan 999999.")
	}
}