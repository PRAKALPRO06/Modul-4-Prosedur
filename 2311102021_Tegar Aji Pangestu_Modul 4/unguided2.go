package main
import "fmt"

//Tegar Aji Pangestu

func hitungSkor(jumlah_soal_selesai * int, skor * int) bool {
	const jumlah_soal int = 8
	var jumlah_soal_selesai_lokal int = jumlah_soal
	var t = [jumlah_soal] int {}
	skor_lokal := 0
	for i := 0; i<jumlah_soal; i++ {
		fmt.Scan(&t[i])
		if t[i] != 301 {
			skor_lokal += t[i]
		} else {
			jumlah_soal_selesai_lokal --
		}
	}
	if jumlah_soal_selesai_lokal > *jumlah_soal_selesai {
		*skor = skor_lokal
		*jumlah_soal_selesai = jumlah_soal_selesai_lokal
		return true
	} else if jumlah_soal_selesai_lokal == *jumlah_soal_selesai {
		if skor_lokal > *skor {
			*skor = skor_lokal
			*jumlah_soal_selesai = jumlah_soal_selesai_lokal
			return true
		}
	}
	return false
}

func main() {
	var jumlah_soal_selesai, skor int
	var nama_pemenang, nama string
	for {
		fmt.Scan(&nama)
		if nama == "Selesai" { break }
		if hitungSkor(&jumlah_soal_selesai, &skor) {
			nama_pemenang = nama
		}
	}
	fmt.Print("\n")
	fmt.Println(nama_pemenang, jumlah_soal_selesai, skor)
}