package main

import (
	"fmt"
	"strings"
)

func hitungSkor(namaPemain *string, jumlahSoal *int, totalSkor *int) {
	var waktuUjian [8]int
	*jumlahSoal = 0
	*totalSkor = 0

	fmt.Print("Masukkan nama peserta: ")
	fmt.Scan(namaPemain)
	if strings.ToLower(*namaPemain) == "selesai" {
		return
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("Masukkan waktu ujian ke-%d: ", i+1)
		fmt.Scan(&waktuUjian[i])
		if waktuUjian[i] != 301 {
			(*jumlahSoal)++
			*totalSkor += waktuUjian[i]
		}
	}
}

func main() {
	var pemenang string
	var soalTertinggi, waktuTerendah int

	for {
		var namaPeserta string
		var jumlahSoal, skorTotal int

		hitungSkor(&namaPeserta, &jumlahSoal, &skorTotal)

		// Cek apakah nama peserta adalah "selesai"
		if strings.ToLower(namaPeserta) == "selesai" {
			break
		}

		fmt.Printf("Peserta: %s, Jumlah Soal: %d, Skor: %d\n", namaPeserta, jumlahSoal, skorTotal)

		if jumlahSoal > soalTertinggi || (jumlahSoal == soalTertinggi && skorTotal < waktuTerendah) {
			pemenang = namaPeserta
			soalTertinggi = jumlahSoal
			waktuTerendah = skorTotal
		}
	}

	fmt.Printf("Pemenang: %s, Jumlah Soal: %d, Skor: %d\n", pemenang, soalTertinggi, waktuTerendah)
}
