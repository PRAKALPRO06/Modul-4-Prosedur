package main

import "fmt"

func hitungSkor(waktu [8]int) (soal int, total int) {
	for i := 0; i < 8; i++ {
		if waktu[i] <= 300 {
			soal = soal + 1
			total = total + waktu[i]
		}
	}
	return soal, total
}

func main() {
	var nama string
	var waktu [8]int
	var pemenangNama string
	var pemenangSoal int = 0
	var pemenangWaktu int = 0

	for {
		fmt.Scanln(&nama, &waktu[0], &waktu[1], &waktu[2], &waktu[3], &waktu[4], &waktu[5], &waktu[6], &waktu[7])

		if nama == "Selesai" {
			break
		}

		soal, total := hitungSkor(waktu)

		if soal > pemenangSoal || (soal == pemenangSoal && total < pemenangWaktu) {
			pemenangNama = nama
			pemenangSoal = soal
			pemenangWaktu = total
		}
	}

	fmt.Println(pemenangNama, pemenangSoal, pemenangWaktu)
}
