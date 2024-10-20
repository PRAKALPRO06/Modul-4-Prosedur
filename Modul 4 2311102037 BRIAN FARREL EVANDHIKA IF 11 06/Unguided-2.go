//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// hitungSkor menghitung jumlah soal yang diselesaikan dan total waktu yang dibutuhkan.
func hitungSkor(waktu []int) (int, int) {
	const waktuMaks = 301 // Waktu maksimum yang diperbolehkan dalam menit
	soalTerselesaikan := 0
	totalWaktu := 0

	for _, w := range waktu {
		if w < waktuMaks { // hanya soal yang diselesaikan dalam waktu kurang dari 301 menit yang dihitung
			soalTerselesaikan++
			totalWaktu += w
		}
	}
	return soalTerselesaikan, totalWaktu
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var namaPemenang string
	var maxSoal, minTotalWaktu int

	// Inisialisasi nilai pemenang
	maxSoal = 0
	minTotalWaktu = 1000000 // Angka yang sangat besar untuk perbandingan

	fmt.Println("Masukkan data peserta (akhiri dengan 'Selesai'): ")
	for scanner.Scan() {
		input := scanner.Text()

		// Memeriksa akhir dari input
		if strings.ToLower(input) == "selesai" {
			break
		}

		// Memisahkan input menjadi bagian-bagian
		parts := strings.Fields(input)

		// Mendapatkan nama peserta
		nama := parts[0]

		// Mengonversi string waktu ke dalam bentuk integer
		var waktu []int
		for i := 1; i < len(parts); i++ {
			w, _ := strconv.Atoi(parts[i])
			waktu = append(waktu, w)
		}

		// Menghitung jumlah soal yang diselesaikan dan total waktu
		soalTerselesaikan, totalWaktu := hitungSkor(waktu)

		// Menentukan apakah peserta ini adalah pemenang baru
		if soalTerselesaikan > maxSoal || (soalTerselesaikan == maxSoal && totalWaktu < minTotalWaktu) {
			namaPemenang = nama
			maxSoal = soalTerselesaikan
			minTotalWaktu = totalWaktu
		}
	}

	// Menampilkan hasil
	fmt.Printf("%s %d %d\n", namaPemenang, maxSoal, minTotalWaktu)
}
