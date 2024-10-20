//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import (
	"fmt"
)

// cetakDeret mencetak setiap suku dari deret berdasarkan aturan yang diberikan.
func cetakDeret(n int) {
	for n != 1 {
		// Mencetak suku saat ini
		fmt.Printf("%d ", n)

		// Menghitung suku berikutnya
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	// Mencetak angka 1 sebagai suku terakhir
	fmt.Println(1)
}

func main() {
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai awal n: ")
	fmt.Scan(&n)

	// Memastikan n valid (kurang dari 1.000.000)
	if n < 1 || n >= 1000000 {
		fmt.Println("Nilai n harus lebih dari 0 dan kurang dari 1.000.000")
		return
	}

	// Memanggil prosedur cetakDeret untuk mencetak deret
	cetakDeret(n)
}
