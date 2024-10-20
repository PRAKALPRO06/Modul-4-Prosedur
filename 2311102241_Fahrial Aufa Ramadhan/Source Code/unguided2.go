package main

import (
	"fmt"
	"strings"
)

func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, t := range waktu {
		if t < 301 {
			*soal++
			*skor += t
		}
	}
}

func main() {
	var pemenang string
	var soalMax, skorMin int
	soalMax = 0
	skorMin = 999999

	for {
		var nama string
		var waktu [8]int

		fmt.Print("Masukkan nama peserta dan waktu penyelesaian soal (ketik 'Selesai' untuk berhenti): ")
		fmt.Scan(&nama)

		if strings.ToLower(nama) == "selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, skor int
		hitungSkor(waktu[:], &soal, &skor)

		if soal > soalMax || (soal == soalMax && skor < skorMin) {
			pemenang = nama
			soalMax = soal
			skorMin = skor
		}
	}

	if pemenang != "" {
		fmt.Printf("Pemenang: %s dengan %d soal diselesaikan dan total waktu %d menit\n", pemenang, soalMax, skorMin)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
