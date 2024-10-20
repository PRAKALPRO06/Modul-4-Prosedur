package main

import (
	"fmt"
	"strings"
)

// Fungsi hitungSkor untuk menghitung jumlah soal yang diselesaikan dan total waktu
func hitungSkor(soal []int) (jumlahSoal int, totalWaktu int) {
	for _, waktu := range soal {
		if waktu <= 300 { // Soal dianggap selesai jika waktu <= 300 menit
			jumlahSoal++
			totalWaktu += waktu
		}
	}
	return
}

func main() {
	var peserta string
	var jumlahSoal int
	var namaPemenang string
	var maxSoal, minWaktu int


	for {
		// Input nama peserta
		fmt.Print("Masukkan nama peserta: ")
		fmt.Scan(&peserta)
		if strings.ToLower(peserta) == "selesai" {
			break
		}

		// Input jumlah soal yang diselesaikan oleh peserta
		fmt.Printf("Masukkan jumlah soal yang diselesaikan oleh %s: ", peserta)
		fmt.Scan(&jumlahSoal)

		// Input waktu untuk setiap soal
		soal := make([]int, jumlahSoal) // membuat array dinamis untuk waktu soal
		fmt.Printf("Masukkan waktu penyelesaian %d soal (dalam menit) untuk %s (dipisahkan dengan spasi):\n", jumlahSoal, peserta)
		for i := 0; i < jumlahSoal; i++ {
			fmt.Scan(&soal[i])
		}

		// Hitung jumlah soal yang diselesaikan dan total waktu
		solvedSoal, totalWaktu := hitungSkor(soal)

		// Tentukan pemenang
		if solvedSoal > maxSoal || (solvedSoal == maxSoal && totalWaktu < minWaktu) {
			namaPemenang = peserta
			maxSoal = solvedSoal
			minWaktu = totalWaktu
		}
	}

	// Output hasil pemenang
	fmt.Printf("Pemenangnya adalah: %s dengan %d soal selesai dan total waktu %d menit\n", namaPemenang, maxSoal, minWaktu)
}
