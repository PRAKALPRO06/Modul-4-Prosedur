package main

import "fmt"

func hitungSukuBerikutnya(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func cetakDeret(n int) {
	fmt.Printf("%d", n)

	for n != 1 {
		n = hitungSukuBerikutnya(n)
		fmt.Printf(" %d", n)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 || n >= 1000000 {
		fmt.Println("Masukkan harus bilangan positif dan kurang dari 1000000")
		return
	}

	cetakDeret(n)
}
