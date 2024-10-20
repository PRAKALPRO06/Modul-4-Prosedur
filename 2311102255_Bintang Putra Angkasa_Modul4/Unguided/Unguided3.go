package main

import "fmt"

// Prosedur untuk mencetak deret sesuai dengan aturan yang diberikan
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Printf("%d\n", n) // Cetak 1 di akhir
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif kurang dari 1000000: ")
	fmt.Scan(&n)
	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Masukkan bilangan yang valid (1 < n < 1000000).")
	}
}
