//WISNU RANANTA RADITYA PUTRA (2311102013) IF-11-06

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hitungSkor(jmlSoal *int, totalSkor *int, waktuPeserta_2311102013 []int) {
	*jmlSoal = 0
	*totalSkor = 0
	for _, waktu := range waktuPeserta_2311102013 {
		if waktu <= 300 {
			*totalSkor += waktu
			*jmlSoal += 1
		}
	}
}
func main() {
	var nama_2311102013 string
	var waktuPeserta_2311102013 []int
	var pemenang_2311102013 string
	var maxSoal, minWaktu int
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "Selesai" {
			break
		}
		input := strings.Fields(line)
		nama_2311102013 = input[0]
		waktuPeserta_2311102013 = make([]int, 0)
		for _, w := range input[1:] {
			waktu, err := strconv.Atoi(w)
			if err == nil {
				waktuPeserta_2311102013 = append(waktuPeserta_2311102013, waktu)
			}
		}
		var jmlSoal, totalSkor int
		hitungSkor(&jmlSoal, &totalSkor, waktuPeserta_2311102013)
		if jmlSoal > maxSoal || (jmlSoal == maxSoal &&
			totalSkor < minWaktu) {
			pemenang_2311102013 = nama_2311102013
			maxSoal = jmlSoal
			minWaktu = totalSkor
		}
	}
	if pemenang_2311102013 != "" {
		fmt.Printf("%s %d %d\n", pemenang_2311102013, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}