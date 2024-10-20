package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, w := range waktu {
		if w <= 300 {
			*soal++
			*skor += w
		}
	}
}

func main() {
	var pemenang string
	var maxSoal, minSkor int
	maxSoal = -1
	minSkor = math.MaxInt32
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Masukkan nama dan waktu :")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "selesai" {
			break
		}

		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		if len(parts) != 9 {
			fmt.Println("Input tidak valid.")
			continue
		}

		nama := parts[0]
		var waktu [8]int

		for i := 0; i < 8; i++ {
			fmt.Sscanf(parts[i+1], "%d", &waktu[i])
		}

		var soal, skor int
		hitungSkor(waktu[:], &soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
		}
	}

	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}

// 2311102266 - hanip 