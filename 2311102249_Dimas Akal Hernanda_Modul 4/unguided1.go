package main

import "fmt"

// Fungsi untuk menghitung faktorial
func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int, hasil *int) {
	var nfaktorial, nrfaktorial int
	faktorial(n, &nfaktorial)
	faktorial(r, &nrfaktorial)
	faktorial(n-r, &nrfaktorial)
	*hasil = nfaktorial / nrfaktorial
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int, hasil *int) {
	var nfaktorial, rfaktorial, nrfaktorial int
	faktorial(n, &nfaktorial)
	faktorial(r, &rfaktorial)
	faktorial(n-r, &nrfaktorial)
	*hasil = nfaktorial / (rfaktorial * nrfaktorial)
}

// Fungsi utama
func main() {
	var a, b, c, d, p1, p2, c1, c2 int
	fmt.Print("Masukkan 4 bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	// Menghitung permutasi dan kombinasi untuk a dan c
	permutasi(a, c, &p1)
	kombinasi(a, c, &c1)
	fmt.Printf("%d %d\n", p1, c1)

	// Menghitung permutasi dan kombinasi untuk b dan d
	permutasi(b, d, &p2)
	kombinasi(b, d, &c2)
	fmt.Printf("%d %d\n", p2, c2)
}
