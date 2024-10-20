package main

import "fmt"

func deretBilangan(n  int)  {
	hasil := n

	for hasil != 1 {
		fmt.Print(hasil, " ")

		if hasil % 2 == 0 {
			hasil = hasil / 2
		} else {
			hasil = (3 * hasil) + 1
		}
	}
	fmt.Println(hasil)
}

func main()  {
	var n int

	fmt.Scan(&n)

	deretBilangan(n)
}