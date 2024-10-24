package main

import "fmt"

func cetakPesan(M string, flag int) {
	var jenis string = ""
	if flag == 0 {
		jenis = "error"
	} else if flag == 1 {
		jenis = "warning"
	} else if flag == 2 {
		jenis = "informasi"
	}
	fmt.Print(M, jenis)
}

func main() {
	var bilangan int
	var pesan string

	fmt.Scan(&bilangan)
	fmt.Scan(&pesan)

	cetakPesan(pesan, bilangan)
}
