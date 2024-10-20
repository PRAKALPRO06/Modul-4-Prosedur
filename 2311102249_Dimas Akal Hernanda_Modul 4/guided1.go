package main

import (
	"fmt"
)

func main() {
	var bilangan int
	var pesan string
	fmt.Scan(&bilangan, &pesan)
	cetakPesan(pesan, bilangan)
}

func cetakPesan(x string, flag int) {
	var jenis string = ""
	if flag == 0 {
		jenis = "Eror Alert!!"
	} else if flag == 1 {
		jenis = "Warning"
	} else if flag == 2 {
		jenis = "Informasi"
	}

	fmt.Println(x, jenis)
}
