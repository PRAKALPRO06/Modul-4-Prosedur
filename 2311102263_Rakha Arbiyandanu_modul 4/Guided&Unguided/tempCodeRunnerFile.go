package main

import "fmt"

// Fungsi f1 mengembalikan hasil bertipe float64
func f1(x, y int) float64 {
	var hasil float64
	hasil = float64(2*x) - 0.5*float64(y) + 3.0
	return hasil
}

// Fungsi f2 menggunakan pointer untuk menyimpan hasil di 'hasil'
func f2(x, y int, hasil *float64) {
	*hasil = float64(2*x) - 0.5*float64(y) + 3.0
}

func main() {
	var a, b int
	var c float64

	fmt.Print("Masukkan dua bilangan bulat: ")
	fmt.Scan(&a, &b)

	// Memanggil f2 untuk menyimpan hasil di c
	f2(a, b, &c)
	fmt.Println("Hasil dari f2 (disimpan di c):", c)

	// Memanggil f1 dan menyimpan hasil di resultF1
	resultF1 := f1(b, a)
	fmt.Println("Hasil dari f1:", resultF1)
}
