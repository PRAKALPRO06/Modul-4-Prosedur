// Meutya Azzahra Efendi
// 2311102166
// IF1106

package main

import (
    "fmt"
)

// Struct untuk menyimpan data peserta
type Peserta struct {
    nama      string
    waktu     [8]int
    soalBenar int
    totalWaktu int
}

// Prosedur untuk input jumlah peserta
func inputJumlahPeserta() int {
    var n int
    fmt.Print("Jumlah Peserta: ")
    fmt.Scan(&n)
    return n
}

// Prosedur untuk input data peserta
func inputDataPeserta(nomorPeserta int) Peserta {
    var peserta Peserta
    
    // Input nama peserta
    fmt.Printf("\nNama peserta %d: ", nomorPeserta+1)
    fmt.Scan(&peserta.nama)
    
    // Input waktu pengerjaan
    fmt.Print("Waktu Pengerjaan Soal (8 soal): ")
    for j := 0; j < 8; j++ {
        fmt.Scan(&peserta.waktu[j])
    }
    
    return peserta
}

// Prosedur untuk menghitung total soal yang dikerjakan dan total waktu
func hitungSkor(waktu [8]int, soal *int, totalWaktu *int) {
    *soal = 0
    *totalWaktu = 0
    for i := 0; i < 8; i++ {
        if waktu[i] <= 300 { // jika waktu pengerjaan kurang dari 300 menit, soal selesai
            *soal++
            *totalWaktu += waktu[i] // hanya tambahkan waktu soal yang selesai
        }
    }
}

// Prosedur untuk menentukan pemenang
func tentukanPemenang(peserta Peserta, pemenangSekarang *Peserta) {
    if peserta.soalBenar > pemenangSekarang.soalBenar || 
       (peserta.soalBenar == pemenangSekarang.soalBenar && peserta.totalWaktu < pemenangSekarang.totalWaktu) {
        *pemenangSekarang = peserta
    }
}

// Prosedur untuk menampilkan hasil akhir
func tampilkanHasil(pemenang Peserta) {
    fmt.Printf("\nNama pemenang: %s\n", pemenang.nama)
    fmt.Printf("Jumlah soal yang selesai: %d\n", pemenang.soalBenar)
    fmt.Printf("Total waktu yang dihabiskan: %d menit\n", pemenang.totalWaktu)
}

func main() {
    // Input jumlah peserta
    n := inputJumlahPeserta()
    
    // Inisialisasi data pemenang sementara
    pemenang := Peserta{
        soalBenar: -1,
        totalWaktu: 1000,
    }
    
    // Proses setiap peserta
    for i := 0; i < n; i++ {
        // Input data peserta
        peserta := inputDataPeserta(i)
        
        // Hitung skor peserta
        hitungSkor(peserta.waktu, &peserta.soalBenar, &peserta.totalWaktu)
        
        // Tentukan pemenang sementara
        tentukanPemenang(peserta, &pemenang)
    }
    
    // Tampilkan hasil akhir
    tampilkanHasil(pemenang)
}