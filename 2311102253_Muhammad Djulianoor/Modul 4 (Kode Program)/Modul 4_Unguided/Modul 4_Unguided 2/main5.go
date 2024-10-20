package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Masukkan jumlah peserta: ")
	fmt.Scanln(&n)

	var pemenang string
	maxSoal := 0
	minWaktu := 999999

	for i := 0; i < n; i++ {
		var nama string
		var soal, waktu int
		fmt.Println("Masukkan data peserta: ")
		fmt.Scan(&nama)

		times := make([]int, 8)
		for j := 0; j < 8; j++ {
			fmt.Scan(&times[j])
		}

		hitungSkor(times, &soal, &waktu)

		if soal > maxSoal || (soal == maxSoal && waktu < minWaktu) {
			maxSoal = soal
			minWaktu = waktu
			pemenang = nama
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
}

func hitungSkor(times []int, soal *int, waktu *int) {
	*soal = 0
	*waktu = 0

	for _, t := range times {
		if t <= 300 {
			*soal++
			*waktu += t
		} else {
			*waktu += 301
		}
	}
}
