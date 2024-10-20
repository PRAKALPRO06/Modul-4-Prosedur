// Meutya Azzahra Efendi
// 2311102166
// IF1106

package main

import (
	"fmt"
)

// Prosedur untuk input bilangan
func inputBilangan() int {
	var n int
	fmt.Print("Bilangan: ")
	fmt.Scan(&n)
	return n
}

// Prosedur untuk validasi input
func validasiInput(n int) bool {
	if n <= 0 || n >= 1000000 {
		fmt.Println("Nilai harus positif dan kurang dari 1000000.")
		return false
	}
	return true
}

// Prosedur untuk menghitung bilangan berikutnya dalam deret
func hitungBilanganBerikutnya(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

// Prosedur untuk menyimpan deret dalam slice
func buatDeret(n int) []int {
	deret := make([]int, 0)
	deret = append(deret, n)

	for n != 1 {
		n = hitungBilanganBerikutnya(n)
		deret = append(deret, n)
	}

	return deret
}

// Prosedur untuk mencetak deret
func cetakDeret(deret []int) {
	fmt.Print("Suku dan deret: ")
	for i, nilai := range deret {
		if i == len(deret)-1 {
			fmt.Printf("%d", nilai)
		} else {
			fmt.Printf("%d ", nilai)
		}
	}
	fmt.Println()
}

// Prosedur untuk analisis deret
func analisisDeret(deret []int) {
	panjangDeret := len(deret)
	nilaiTerbesar := deret[0]

	for _, nilai := range deret {
		if nilai > nilaiTerbesar {
			nilaiTerbesar = nilai
		}
	}

	fmt.Printf("\nAnalisis Deret:\n")
	fmt.Printf("Panjang deret: %d\n", panjangDeret)
	fmt.Printf("Nilai terbesar dalam deret: %d\n", nilaiTerbesar)
}

func main() {
	// Input bilangan
	n := inputBilangan()

	// Validasi input
	if validasiInput(n) {
		// Buat deret
		deret := buatDeret(n)

		// Cetak deret
		cetakDeret(deret)

		// Tampilkan analisis deret
		analisisDeret(deret)
	}
}