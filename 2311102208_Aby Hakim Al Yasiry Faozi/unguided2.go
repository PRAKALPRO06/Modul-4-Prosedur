package main

import (
	"fmt"
)

func hitungSkor(soal []int, skor *int) int {
	totalSkor := 0
	totalSoal := 0
	for _, waktu := range soal {
		if waktu < 301 {
			totalSkor += waktu
			totalSoal++
		}
	}
	*skor = totalSkor
	return totalSoal
}

func main() {
	var peserta1, peserta2 [8]int
	var skor1, skor2 int

	fmt.Println("Masukkan waktu pengerjaan soal Astuti:")
	fmt.Scanln(&peserta1[0], &peserta1[1], &peserta1[2], &peserta1[3], &peserta1[4], &peserta1[5], &peserta1[6], &peserta1[7])

	fmt.Println("Masukkan waktu pengerjaan soal Bertha:")
	fmt.Scanln(&peserta2[0], &peserta2[1], &peserta2[2], &peserta2[3], &peserta2[4], &peserta2[5], &peserta2[6], &peserta2[7])

	soalTerselesaikan1 := hitungSkor(peserta1[:], &skor1)
	soalTerselesaikan2 := hitungSkor(peserta2[:], &skor2)

	fmt.Println("Keluaran:")

	if soalTerselesaikan1 > soalTerselesaikan2 {
		fmt.Printf("Pemenang: Astuti\nJumlah soal diselesaikan: %d\nTotal waktu: %d\n", soalTerselesaikan1, skor1)
	} else if soalTerselesaikan2 > soalTerselesaikan1 {
		fmt.Printf("Pemenang: Bertha\nJumlah soal diselesaikan: %d\nTotal waktu: %d\n", soalTerselesaikan2, skor2)
	} else {
		if skor1 < skor2 {
			fmt.Printf("Pemenang: Astuti\nJumlah soal diselesaikan: %d\nTotal waktu: %d\n", soalTerselesaikan1, skor1)
		} else if skor2 < skor1 {
			fmt.Printf("Pemenang: Bertha\nJumlah soal diselesaikan: %d\nTotal waktu: %d\n", soalTerselesaikan2, skor2)
		} else {
			fmt.Println("Hasil seri!")
		}
	}
}
