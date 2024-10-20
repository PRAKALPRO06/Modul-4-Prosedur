package main

import "fmt"

func Cetakderet(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal deret: ")
	fmt.Scan(&n)

	Cetakderet(n)
}
