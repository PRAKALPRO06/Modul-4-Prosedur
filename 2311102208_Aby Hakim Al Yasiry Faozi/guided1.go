package main

import "fmt"

func main() {
	var bilangan int
	var pesan string

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	fmt.Print("Masukkan pesan: ")
	fmt.Scanln(&pesan)
	fmt.Scanln(&pesan)

	cetakPesan(pesan, bilangan)
}

func cetakPesan(M string, flag int) {
	var jenis string = ""
	switch flag {
	case 0:
		jenis = "eror"
	case 1:
		jenis = "warning"
	case 2:
		jenis = "informasi"
	default:
		jenis = "tidak dikenal"
	}
	fmt.Println(M, jenis)
}
