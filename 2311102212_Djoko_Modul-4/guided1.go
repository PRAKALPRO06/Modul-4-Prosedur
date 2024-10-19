package main

import (
	"fmt"
)


func cetakPesan(pesan string, flag int) {
	var jenis string
	if flag == 0 {
		jenis = "Error"
	} else if flag == 1 {
		jenis = "Warning"
	} else if flag == 2 {
		jenis = "Informasi"
	}
	fmt.Printf("%s: %s\n", jenis, pesan)
}

func main() {
	var bilangan int
	var pesan string

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	fmt.Print("Masukkan pesan: ")
	fmt.Scan(&pesan)

	cetakPesan(pesan, bilangan)
}
