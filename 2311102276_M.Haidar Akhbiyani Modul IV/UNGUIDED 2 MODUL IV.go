package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hitungSkor(waktu []int) (int, int) {
	const waktuMaks = 301 
	soalTerselesaikan := 0
	totalWaktu := 0

	for _, w := range waktu {
		if w < waktuMaks { 
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

	maxSoal = 0
	minTotalWaktu = 1000000 
	fmt.Println("Masukkan data peserta (akhiri dengan 'Selesai'): ")
	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		parts := strings.Fields(input)

		nama := parts[0]

		var waktu []int
		for i := 1; i < len(parts); i++ {
			w, _ := strconv.Atoi(parts[i])
			waktu = append(waktu, w)
		}

		soalTerselesaikan, totalWaktu := hitungSkor(waktu)

		if soalTerselesaikan > maxSoal || (soalTerselesaikan == maxSoal && totalWaktu < minTotalWaktu) {
			namaPemenang = nama
			maxSoal = soalTerselesaikan
			minTotalWaktu = totalWaktu
		}
	}

	fmt.Printf("%s %d %d\n", namaPemenang, maxSoal, minTotalWaktu)
}