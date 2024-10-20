package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hitungSkor(waktu []int, soal *int, skor *int, waktuTotal *int) {
	*soal = 0
	*skor = 0
	*waktuTotal = 0

	for _, menit := range waktu {
		if menit <= 300 {
			*soal++
			*skor += 300 - menit
			*waktuTotal += menit
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var namaPemenang string
	var soalPemenang, waktuTotalPemenang int
	selesai := false // Variabel untuk menandakan status selesai

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Cek apakah input adalah "Selesai"
		if input == "Selesai" && !selesai {
			selesai = true // Set status selesai
			break
		}

		parts := strings.Fields(input)
		nama := parts[0]
		waktu := make([]int, len(parts)-1)

		for i := 1; i < len(parts); i++ {
			menit, err := strconv.Atoi(parts[i])
			if err == nil {
				waktu[i-1] = menit
			}
		}

		var soal, skor, waktuTotal int
		hitungSkor(waktu, &soal, &skor, &waktuTotal)

		if (namaPemenang == "") || (skor > waktuTotalPemenang) || (skor == waktuTotalPemenang && waktuTotal < waktuTotalPemenang) {
			namaPemenang = nama
			soalPemenang = soal
			waktuTotalPemenang = waktuTotal
		}
	}

	if namaPemenang != "" {
		fmt.Println()
		fmt.Printf("%s %d %d\n", namaPemenang, soalPemenang, waktuTotalPemenang)
	}
}
