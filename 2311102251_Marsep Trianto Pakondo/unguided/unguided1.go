package main

import "fmt"

var a,b,c,d int

// Fungsi untuk menghitung faktorial dari bilangan n
func faktorial(n int, hasil *int) {
	*hasil = 1
	// Loop untuk menghitung faktorial, dimulai dari 1 hingga n
	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

// Fungsi untuk menghitung permutasi P(n, r) = n! / (n-r)!
func permutasi(n, r int)  {
	var faktorN, faktorNR int
	faktorial(n, &faktorN)
	faktorial(n-r, &faktorNR)

	hasil := faktorN / faktorNR

	fmt.Print(hasil)
}

// Fungsi untuk menghitung kombinasi C(n, r) = n! / (r! * (n-r)!)
func kombinasi(n, r int) {
	var faktorN, faktorR, faktorNR int
	faktorial(n, &faktorN)
	faktorial(r, &faktorR)
	faktorial(n-r, &faktorNR)

	hasil := faktorN / (faktorR * faktorNR)

	fmt.Print(hasil)
}

func main() {

	fmt.Print("Masukkan input = ")
	fmt.Scan(&a,&b,&c,&d)

	if a >= c && b >= d {
		// Baris pertama: Permutasi dan Kombinasi a terhadap c
		permutasi(a, c)
		fmt.Print(", ")
		kombinasi(a, c)
		
		fmt.Print("\n")
		// Baris kedua: Permutasi dan Kombinasi b terhadap d
		permutasi(b, d) 
		fmt.Print(", ")
		kombinasi(b, d)
	} else {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d")
	}
}