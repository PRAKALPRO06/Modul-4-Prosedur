package main

import (
	"fmt"
)

// Prosedur untuk menghitung jumlah soal yang diselesaikan dan total skor
func hitungSkor(waktu [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for i := 0; i < 8; i++ {
		if waktu[i] <= 300 { // Jika soal diselesaikan dalam waktu <= 300 menit
			*soal++
			*skor += waktu[i]
		} else {
			*skor += 301 // Jika soal tidak diselesaikan, tambahkan 301 menit
		}
	}
}

func main() {
	var jumlahPeserta int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	// Variabel untuk menyimpan informasi pemenang
	var pemenangNama string
	var pemenangSoal, pemenangSkor int

	// Loop untuk setiap peserta
	for i := 0; i < jumlahPeserta; i++ {
		// Membaca nama peserta
		var namaPeserta string
		fmt.Print("Masukkan nama peserta: ")
		fmt.Scan(&namaPeserta)

		// Membaca waktu pengerjaan 8 soal
		var waktu [8]int
		fmt.Print("Masukkan waktu pengerjaan 8 soal (dalam menit): ")
		for j := 0; j < 8; j++ {
			fmt.Scan(&waktu[j])
		}

		// Variabel untuk menyimpan hasil skor peserta saat ini
		var soal, skor int

		// Menghitung jumlah soal yang diselesaikan dan total skor
		hitungSkor(waktu, &soal, &skor)

		// Menentukan pemenang
		if soal > pemenangSoal || (soal == pemenangSoal && skor < pemenangSkor) {
			pemenangNama = namaPeserta
			pemenangSoal = soal
			pemenangSkor = skor
		}
	}

	// Output pemenang
	fmt.Printf("Pemenangnya adalah %s yang menyelesaikan %d soal dengan total waktu %d menit.\n", pemenangNama, pemenangSoal, pemenangSkor)
}
