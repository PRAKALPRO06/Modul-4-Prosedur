package main

import (
	"fmt"
	"strings"
)

func hitungSkor(waktu []int, totalSoal *int, totalSkor *int) {
	*totalSoal = 0
	*totalSkor = 0
	for _, t := range waktu {
		if t < 301 { // Waktu penyelesaian kurang dari 301 menit
			*totalSoal++
			*totalSkor += t
		}
	}
}

func main() {
	var pemenang string
	var maxSoal, minSkor int
	maxSoal = 0
	minSkor = 999999

	for {
		var nama string
		var waktu [8]int

		fmt.Print("Masukkan nama peserta (ketik 'Selesai' untuk menghentikan pendaftaran): ")
		fmt.Scan(&nama)

		if strings.ToLower(nama) == "selesai" {
			break
		}

		fmt.Println("Masukkan waktu penyelesaian untuk 8 soal (dalam menit, satu per satu):")
		for i := 0; i < 8; i++ {
			_, err := fmt.Scan(&waktu[i])
			if err != nil {
				fmt.Println("Terjadi kesalahan dalam membaca input. Pastikan Anda memasukkan angka yang benar.")
				i-- // Mengulangi pembacaan waktu untuk soal yang sama
			}
		}

		var soal, skor int
		hitungSkor(waktu[:], &soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
		}
	}

	if pemenang != "" {
		fmt.Printf("Pemenang adalah: %s\n", pemenang)
		fmt.Printf("Jumlah soal yang diselesaikan: %d\n", maxSoal)
		fmt.Printf("Total waktu penyelesaian: %d menit\n", minSkor)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
