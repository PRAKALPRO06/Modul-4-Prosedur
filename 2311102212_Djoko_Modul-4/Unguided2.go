package main

import (
	"fmt"
)

func main() {
	// Kamus
	var nilai [2][8]int
	var nama1, nama2 string
	var i, j int

	// Algoritma
	for i = 0; i < 2; i++ {
		if i == 0 {
			fmt.Print("Masukkan nama peserta 1: ")
			fmt.Scan(&nama1)
		} else {
			fmt.Print("Masukkan nama peserta 2: ")
			fmt.Scan(&nama2)
		}

		for j = 0; j < 8; j++ {
			fmt.Printf("Masukkan nilai untuk peserta %d, soal ke-%d: ", i+1, j+1)
			fmt.Scan(&nilai[i][j])
			if nilai[i][j] == 0 {
				nilai[i][j] = 301
			}
		}
	}

	// Menghitung skor
	hitungSkor(nilai, 2, 8, nama1, nama2)
}

func hitungSkor(arr [2][8]int, jumlahOrang, banyakSkor int, peserta1, peserta2 string) {
	// Kamus
	var hasilSkor1, hasilSkor2 int
	var jumlahSoalTerjawab1, jumlahSoalTerjawab2 int

	// Algoritma
	for i := 0; i < jumlahOrang; i++ {
		for j := 0; j < banyakSkor; j++ {
			if i == 0 {
				hasilSkor1 += arr[i][j]
				if arr[i][j] != 301 {
					jumlahSoalTerjawab1++
				}
			} else {
				hasilSkor2 += arr[i][j]
				if arr[i][j] != 301 {
					jumlahSoalTerjawab2++
				}
			}
		}
	}

	// Menentukan pemenang
	if jumlahSoalTerjawab1 == jumlahSoalTerjawab2 {
		if hasilSkor1 > hasilSkor2 {
			hasilSkor1 -= 301 * (banyakSkor - jumlahSoalTerjawab1)
			fmt.Printf("Pemenang: %s, Jumlah soal terjawab: %d, Total skor: %d\n", peserta1, jumlahSoalTerjawab1, hasilSkor1)
		} else {
			hasilSkor2 -= 301 * (banyakSkor - jumlahSoalTerjawab2)
			fmt.Printf("Pemenang: %s, Jumlah soal terjawab: %d, Total skor: %d\n", peserta2, jumlahSoalTerjawab2, hasilSkor2)
		}
	} else if jumlahSoalTerjawab1 > jumlahSoalTerjawab2 {
		hasilSkor1 -= 301 * (banyakSkor - jumlahSoalTerjawab1)
		fmt.Printf("Pemenang: %s, Jumlah soal terjawab: %d, Total skor: %d\n", peserta1, jumlahSoalTerjawab1, hasilSkor1)
	} else {
		hasilSkor2 -= 301 * (banyakSkor - jumlahSoalTerjawab2)
		fmt.Printf("Pemenang: %s, Jumlah soal terjawab: %d, Total skor: %d\n", peserta2, jumlahSoalTerjawab2, hasilSkor2)
	}
}
