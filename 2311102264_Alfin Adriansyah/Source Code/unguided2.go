package main

import (
	"fmt"
)

func hitungSkor(waktu []int, jumlahSoal *int, totalSkor *int) {
	*jumlahSoal = 0
	*totalSkor = 0
	for _, t := range waktu {
		if t < 301 {
			*jumlahSoal++
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
		var namaPeserta string
		var waktuPenyelesaian [8]int

		fmt.Print("Masukkan nama dan waktu penyelesaian ('Selesai' untuk berhenti): ")
		fmt.Scan(&namaPeserta)

		if namaPeserta == "Selesai" || namaPeserta == "selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktuPenyelesaian[i])
		}

		var jumlahSoal, totalWaktu int
		hitungSkor(waktuPenyelesaian[:], &jumlahSoal, &totalWaktu)

		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minSkor) {
			pemenang = namaPeserta
			maxSoal = jumlahSoal
			minSkor = totalWaktu
		}
	}

	if pemenang != "" {
		fmt.Printf("Pemenang: %s dengan %d soal diselesaikan dan total waktu %d menit\n", pemenang, maxSoal, minSkor)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
