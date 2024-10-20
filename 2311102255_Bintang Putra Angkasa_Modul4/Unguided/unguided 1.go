package main

import "fmt"

// Fungsi untuk menghitung faktorial
func faktorial(f int, hasil *int) {
	*hasil = 1
	for i := 1; i <= f; i++ {
		*hasil *= i
	}
}

// Fungsi untuk memeriksa syarat
func ceksyarat(a, b, c, d int, syarat *bool) {
	*syarat = a >= c && b >= d
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int, hasil *int) {
	var faktorialN, faktorialNR int
	faktorial(n, &faktorialN)
	faktorial(n-r, &faktorialNR)
	*hasil = faktorialN / faktorialNR
}

func kombinasi(n, r int, hasil *int) {
	var faktorialN, faktorialR, faktorialNR int
	faktorial(n, &faktorialN)
	faktorial(n-r, &faktorialNR)
	faktorial(r, &faktorialR)
	*hasil = faktorialN / (faktorialR * faktorialNR)
}

func main() {
	var a, b, c, d int
	var syarat bool
	var hasil int

	fmt.Print("Masukkan nilai a, b, c, dan d: ")
	fmt.Scan(&a, &b, &c, &d)

	// Memeriksa syarat
	ceksyarat(a, b, c, d, &syarat)

	if syarat {
		permutasi(a, c, &hasil)
		fmt.Printf("hasil permutasi  %d dan %d adalah %d \n", a, c, hasil)
		kombinasi(a, c, &hasil)
		fmt.Printf("hasil kombinasi  %d dan %d adalah %d \n", a, c, hasil)
		permutasi(b, d, &hasil)
		fmt.Printf("hasil permutasi  %d dan %d adalah %d \n", b, d, hasil)
		kombinasi(b, d, &hasil)
		fmt.Printf("hasil kombinasi  %d dan %d adalah %d \n", b, d, hasil)

	} else {
		fmt.Println("Inputan yang anda masukkan salah")
	}
}
