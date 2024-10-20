package main
import "fmt"
//Tegar Aji Pangestu
var a, b, c, d int

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func hitungPermutasi(n, r int) {
	hasil := faktorial(n) / faktorial(n-r)
	fmt.Printf("Permutasi(%d, %d) = %d\n", n, r, hasil)
}

func hitungKombinasi(n, r int) {
	hasil := faktorial(n) / (faktorial(r) * faktorial(n-r))
	fmt.Printf("Kombinasi(%d, %d) = %d\n", n, r, hasil)
}
func main() {
	fmt.Print("Masukkan input (a, b, c, d) = ")
	fmt.Scan(&a, &b, &c, &d)
	if a >= c && b >= d {
		hitungPermutasi(a, c)
		hitungKombinasi(a, c)
		hitungPermutasi(b, d)
		hitungKombinasi(b, d)
	} else {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d")
	}
}
