package main

import (
	"fmt"
)

func CetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var input int
	fmt.Scan(&input)

	if input > 0 && input < 1000000 {
		CetakDeret(input)
	} else {
		fmt.Println("Masukan harus bilangan positif kurang dari 1000000")
	}
}