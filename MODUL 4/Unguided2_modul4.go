package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, t := range waktu {
		if t <= 300 {
			*soal++
			*skor += t
		}
	}
}

func bacaInput(scanner *bufio.Scanner) (string, []int) {
	line := scanner.Text()
	fields := strings.Fields(line)
	nama := fields[0]
	waktu := make([]int, 8)
	for i := 0; i < 8; i++ {
		waktu[i], _ = strconv.Atoi(fields[i+1])
	}
	return nama, waktu
}

func updatePemenang(nama string, soal int, skor int, pemenang *string, maxSoal *int, minSkor *int) {
	if soal > *maxSoal || (soal == *maxSoal && skor < *minSkor) {
		*pemenang = nama
		*maxSoal = soal
		*minSkor = skor
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pemenang string
	var maxSoal, minSkor int
	minSkor = 9999

	for scanner.Scan() {
		if scanner.Text() == "Selesai" {
			break
		}

		nama, waktu := bacaInput(scanner)
		var soal, skor int
		hitungSkor(waktu, &soal, &skor)
		updatePemenang(nama, soal, skor, &pemenang, &maxSoal, &minSkor)
	}

	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
}
