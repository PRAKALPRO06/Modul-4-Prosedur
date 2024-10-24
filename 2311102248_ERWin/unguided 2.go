package main

import "fmt"

func hitungSkor(soalSelesai, skor *int) {
	var waktu int
	*soalSelesai = 0
	*skor = 0
	soal := 8

	for i := 1; i <= soal; i++ {
		fmt.Printf("masukkan waktu selesai soal %v : ", i)
		fmt.Scanln(&waktu)

		if waktu < 301 {
			*soalSelesai++
			*skor = *skor + waktu
		}
	}
}

func main() {
	var winn48, erwinn_248 string
	var skor, soalSelesai int

	for {
		var temp, temp1 int

		fmt.Print("\nMasukkan nama peserta : ")
		fmt.Scanln(&winn48)

		if winn48 == "selesai" {
			break
		}

		hitungSkor(&temp, &temp1)

		if temp > soalSelesai || temp == soalSelesai && temp1 < skor {
			erwinn_248 = winn48
			soalSelesai = temp
			skor = temp1
		}
	}

	fmt.Printf("\n%v %v %v", erwinn_248, soalSelesai, skor)
}
