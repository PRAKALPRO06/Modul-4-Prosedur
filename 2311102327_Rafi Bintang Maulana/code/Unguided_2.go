package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HitungSkor(waktu []int) (int, int) {
	totalPenyelesaian := 0
	totalWaktu := 0
	for _, w := range waktu {
		if w <= 300 {
			totalPenyelesaian++
			totalWaktu += w
		}
	}
	return totalPenyelesaian, totalWaktu
}

func main() {
	pemindai := bufio.NewScanner(os.Stdin)
	var namaPemenang string
	maksPenyelesaian := 0
	minWaktu := 999999
	adaPesertaMenyelesaikan := false

	for {
		pemindai.Scan()
		masukan := pemindai.Text()

		if masukan == "Selesai" {
			break
		}

		bagian := strings.Fields(masukan)
		nama := bagian[0]
		waktu := make([]int, len(bagian)-1)

		for i := 1; i < len(bagian); i++ {
			fmt.Sscanf(bagian[i], "%d", &waktu[i-1])
		}

		totalPenyelesaian, totalWaktu := HitungSkor(waktu)

		if totalPenyelesaian > 0 {
			adaPesertaMenyelesaikan = true
		}

		if totalPenyelesaian > maksPenyelesaian || (totalPenyelesaian == maksPenyelesaian && totalWaktu < minWaktu) {
			namaPemenang = nama
			maksPenyelesaian = totalPenyelesaian
			minWaktu = totalWaktu
		}
	}

	if adaPesertaMenyelesaikan {
		fmt.Printf("%s %d %d\n", namaPemenang, maksPenyelesaian, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta yang menyelesaikan masalah dalam batas waktu.")
	}
}