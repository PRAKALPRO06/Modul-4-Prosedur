package main

import (
	"fmt"
)

func hitungSkor(soal *int, skor *int, waktu []int) {
    *soal = 0
    *skor = 0
    
    // Menghitung jumlah soal yang diselesaikan dan total waktu
    for _, t := range waktu {
        if t != 301 {  // Jika waktu bukan 301 atau berhasil diselesaikan
            *soal++
            *skor += t
        }
    }
}

func main() {
    var peserta []string
    var waktu [][]int
    
    // Akan membaca input sampai nama "Selesai"
    for {
        var nama string
        fmt.Scan(&nama)
        if nama == "Selesai" {
            break
        }
        
        // Membaca 8 waktu pengerjaan
        var times []int
        for i := 0; i < 8; i++ {
            var t int
            fmt.Scan(&t)
            times = append(times, t)
        }
        
        peserta = append(peserta, nama)
        waktu = append(waktu, times)
    }
    
    // Mencari pemenang
    var maxSoal_226, minWaktu_226 int = -1, 10000
    var pemenang_226 string
    
    for i := 0; i < len(peserta); i++ {
        var jmlSoal, totalWaktu int
        hitungSkor(&jmlSoal, &totalWaktu, waktu[i])

        if jmlSoal > maxSoal_226 || (jmlSoal == maxSoal_226 && totalWaktu < minWaktu_226) {
            maxSoal_226 = jmlSoal
            minWaktu_226 = totalWaktu
            pemenang_226 = peserta[i]
        }
    }
    
    // Cetak hasil
    fmt.Printf("%s %d %d\n", pemenang_226, maxSoal_226, minWaktu_226)
}
