package main

import "fmt"

// Prosedur untuk mencetak deret bilangan
func cetakDeret(n_2217 int) {
	for n_2217 != 1 {
		fmt.Print(n_2217, " ")
		if n_2217%2 == 0 {
			n_2217 = n_2217 / 2
		} else {
			n_2217 = 3*n_2217 + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
