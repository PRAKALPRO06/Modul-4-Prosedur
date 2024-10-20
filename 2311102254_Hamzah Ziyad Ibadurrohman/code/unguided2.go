package main

import "fmt"

var nama string
var soal, skor int

func HitungSkor(jawab [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if jawab[i] != 301 {
			(*soal)++
			*skor += jawab[i]
		}
	}
}

func main() {
	var jawab [8]int
	var nama, pemenang string
	var skormin, soalmaks int

	soalmaks = 0
	skormin = 999999

	for {
		fmt.Print("Nama peserta: ")
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		fmt.Println("Waktu pengerjaan 8 soal: ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&jawab[i])
		}

		var soalSelesai, totalWaktu int
		HitungSkor(jawab, &soalSelesai, &totalWaktu)

		if soalSelesai > soalmaks || (soalSelesai == soalmaks && totalWaktu < skormin) {
			pemenang = nama
			soalmaks = soalSelesai
			skormin = totalWaktu
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, soalmaks, skormin)
}
