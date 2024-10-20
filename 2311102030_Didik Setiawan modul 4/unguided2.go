package main

import (
	"fmt"
	"strings"
)

func hitungSkor(namaPeserta *string, jumlahSoalDijawab *int, totalWaktu *int) {
	var waktuUjianPerSoal [8]int
	*jumlahSoalDijawab = 0
	*totalWaktu = 0

	fmt.Print("Masukkan nama peserta: ")
	fmt.Scan(namaPeserta)
	if strings.ToLower(*namaPeserta) == "selesai" {
		return
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("Masukkan waktu ujian ke-%d: ", i+1)
		fmt.Scan(&waktuUjianPerSoal[i])
		if waktuUjianPerSoal[i] != 301 {
			(*jumlahSoalDijawab)++
			*totalWaktu += waktuUjianPerSoal[i]
		}
	}
}

func main() {
	var namaPemenang string
	var soalTerbanyak, waktuTercepat int

	for {
		var namaPeserta string
		var jumlahSoalDijawab, totalWaktu int

		hitungSkor(&namaPeserta, &jumlahSoalDijawab, &totalWaktu)

		// Cek apakah nama peserta adalah "selesai"
		if strings.ToLower(namaPeserta) == "selesai" {
			break
		}

		fmt.Printf("Peserta: %s, Jumlah Soal Dijawab: %d, Total Waktu: %d\n", namaPeserta, jumlahSoalDijawab, totalWaktu)

		if jumlahSoalDijawab > soalTerbanyak || (jumlahSoalDijawab == soalTerbanyak && totalWaktu < waktuTercepat) {
			namaPemenang = namaPeserta
			soalTerbanyak = jumlahSoalDijawab
			waktuTercepat = totalWaktu
		}
	}

	fmt.Printf("Pemenang: %s, Jumlah Soal Dijawab: %d, Total Waktu: %d\n", namaPemenang, soalTerbanyak, waktuTercepat)
}
