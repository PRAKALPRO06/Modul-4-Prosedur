//Reza Alvonzo//
package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hitungSkor(jumlahSoal *int, totalSkor *int, waktuPeserta []int) {
	*jumlahSoal = 0
	*totalSkor = 0
	for _, waktu := range waktuPeserta {
		if waktu <= 300 {
			*totalSkor += waktu
			*jumlahSoal += 1
		}
	}
}
func main() {
	var namaPeserta string
	var waktuPeserta []int
	var pemenang string
	var maxSoal, minWaktu int
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "Selesai" {
			break
		}
		input := strings.Fields(line)
		namaPeserta = input[0]
		waktuPeserta = make([]int, 0)
		for _, w := range input[1:] {
			waktu, err := strconv.Atoi(w)
			if err == nil {
				waktuPeserta = append(waktuPeserta, waktu)
			}
		}
		var jumlahSoal, totalSkor int
		hitungSkor(&jumlahSoal, &totalSkor, waktuPeserta)
		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal &&
			totalSkor < minWaktu) {
			pemenang = namaPeserta
			maxSoal = jumlahSoal
			minWaktu = totalSkor
		}
	}
	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
