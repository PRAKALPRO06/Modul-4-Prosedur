package main

import (
	"fmt"
)

func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("input angka : ")
	fmt.Scanln(&a, &b, &c, &d)
	if a >= c && b >= d {
		fmt.Println(permutasi(a, c), kombinasi(a, c))
		fmt.Println(permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Print("MAAF ANDA TIDAK MEMENUHI PERSYARATAN")
	}

}
