package main

import "fmt"

var a, b, c, d int
var hasilPermutasi, hasilKombinasi int

func faktorial(n int, hasil *int) {
	*hasil = 1

	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutasi(n, r int, hasil *int) {
	var faktN, faktNMinusR int
	faktorial(n, &faktN)
	faktorial(n-r, &faktNMinusR)
	*hasil = faktN / faktNMinusR
}

func kombinasi(n, r int, hasil *int) {
	var faktN, faktR, faktNMinusR int
	faktorial(n, &faktN)
	faktorial(r, &faktR)
	faktorial(n-r, &faktNMinusR)
	*hasil = faktN / (faktR * faktNMinusR)
}

func main() {

	fmt.Print("Masukkan input = ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {

		permutasi(a, c, &hasilPermutasi)
		kombinasi(a, c, &hasilKombinasi)
		fmt.Printf("%d, %d\n", hasilPermutasi, hasilKombinasi)

		permutasi(b, d, &hasilPermutasi)
		kombinasi(b, d, &hasilKombinasi)
		fmt.Printf("%d, %d\n", hasilPermutasi, hasilKombinasi)
	} else {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d")
	}
}
