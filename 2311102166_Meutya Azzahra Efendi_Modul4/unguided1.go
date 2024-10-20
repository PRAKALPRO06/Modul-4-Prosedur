// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

var a, b, c, d int

// Fungsi untuk menghitung faktorial dari bilangan n
func faktorial(n int) int {
	hasil := 1
	// Loop untuk menghitung faktorial, dimulai dari 1 hingga n
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r) = n! / (n-r)!
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r) = n! / (r! * (n-r)!)
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {

	fmt.Print("Masukkan input = ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		// Baris pertama: Permutasi dan Kombinasi a terhadap c
		fmt.Printf("%d, %d\n", permutasi(a, c), kombinasi(a, c))

		// Baris kedua: Permutasi dan Kombinasi b terhadap d
		fmt.Printf("%d, %d\n", permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus>=d")
	}
}
