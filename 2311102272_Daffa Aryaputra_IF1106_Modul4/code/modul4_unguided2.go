package main

import (
	"fmt"
	"strings"
)

func hitungSkor(soal *int, skor *int) string {
	var nama string
	var waktu [8]int
	*soal = 0
	*skor = 0

	fmt.Scan(&nama)
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu[i])
		if waktu[i] != 301 {
			(*soal)++
			*skor += waktu[i]
		}
	}
	return nama
}

func main() {
	var pemenang string
	var maxSoal, minWaktu int

	for {
		var soal, skor int
		
		// Panggil fungsi hitungSkor dan ambil nama
		nama := hitungSkor(&soal, &skor)
		
		// Cek apakah nama adalah "selesai"
		if strings.ToLower(nama) == "selesai" {
			break
		}
		
		// Update pemenang jika memenuhi kriteria
		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}
	
	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
}