package main

import (
	"fmt"
)

func cetakDeret(n int) {
	deret := []int{}
	
	deret = append(deret, n)

	for n != 1 {
		if n%2 == 0 { 
			n = n / 2
		} else { 
			n = 3*n + 1
		}
		
		deret = append(deret, n)
	}

	for i, v := range deret {
		if i == len(deret)-1 {
			fmt.Print(v) 
		} else {
			fmt.Print(v, " ") 
		}
	}
	fmt.Println() 
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan positif yang lebih kecil dari 1000000: ")
	fmt.Scan(&n)


	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Input tidak valid. Pastikan bilangan positif dan kurang dari 1000000.")
	}
}
