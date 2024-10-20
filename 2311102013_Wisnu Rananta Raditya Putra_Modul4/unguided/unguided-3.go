//WISNU RANANTA RADITYA PUTRA (2311102013) IF-11-06

package main

import (
	"fmt"
)

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n_2311102013 int
	fmt.Println("Masukkan nilai suku awal (bilangan positif < 1000000):")
	fmt.Scan(&n_2311102013)

	if n_2311102013 <= 0 || n_2311102013 >= 1000000 {
		fmt.Println("Masukan tidak valid, masukkan bilangan positif kurang dari 1000000.")
		return
	}

	cetakDeret(n_2311102013)
}