// 2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import (
	"fmt"
)

// cetakDeret mencetak setiap elemen dari deret berdasarkan aturan yang diberikan.
func cetakDeret(angka int) {
	for angka != 1 {
		// Mencetak elemen saat ini
		fmt.Printf("%d ", angka)

		// Menghitung elemen berikutnya
		if angka%2 == 0 {
			angka = angka / 2
		} else {
			angka = 3*angka + 1
		}
	}
	// Mencetak angka 1 sebagai elemen terakhir
	fmt.Println(1)
}

func main() {
	var nilaiAwal int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&nilaiAwal)

	// Memastikan nilaiAwal valid (kurang dari 1.000.000)
	if nilaiAwal < 1 || nilaiAwal >= 1000000 {
		fmt.Println("Nilai harus lebih dari 0 dan kurang dari 1.000.000")
		return
	}

	// Memanggil prosedur cetakDeret untuk mencetak deret
	cetakDeret(nilaiAwal)
}
