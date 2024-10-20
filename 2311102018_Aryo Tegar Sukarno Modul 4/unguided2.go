//2311102018 Aryo Tegar Sukarno
package main

import (
        "fmt"
        "strconv"
        "strings"
)
type Peserta struct {
        nama  string
        soal  int
        skor  int
}

func hitungSkor(waktu []int) (int, int) {
        skor := 0
        soal := 0
        for _, t := range waktu {
                if t <= 300 {
                        skor += t
                        soal++
                }
        }
        return soal, skor
}

func main() {
        var peserta []Peserta

        
        for {
            var baris string
            fmt.Scanln(&baris)
            if baris == "Selesai" {
                break
            }

        
            data := strings.Split(baris, " ")
            nama := data[0]
            waktu := make([]int, 8)
            for i := 1; i < len(data); i++ {
                waktu[i-1], _ = strconv.Atoi(data[i])
            }

            
            soal, skor := hitungSkor(waktu)
            peserta = append(peserta, Peserta{nama, soal, skor})
        }
        pemenang := peserta[0]
        for _, p := range peserta[1:] {
                if p.soal > pemenang.soal || (p.soal == pemenang.soal && p.skor < pemenang.skor) {
                        pemenang = p
                }
        }
        fmt.Printf("%s %d %d\n", pemenang.nama, pemenang.soal, pemenang.skor)
}