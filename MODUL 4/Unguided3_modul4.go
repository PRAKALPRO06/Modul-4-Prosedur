package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cetakDeret(n int) string {
	var result []string
	result = append(result, strconv.Itoa(n))

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		result = append(result, strconv.Itoa(n))
	}

	return strings.Join(result, " ")
}

func main() {
	var input int
	fmt.Print("Masukan: ")
	fmt.Scan(&input)

	if input <= 0 || input >= 1000000 {
		fmt.Println("Masukan harus berupa bilangan integer positif yang lebih kecil dari 1000000.")
		return
	}

	fmt.Println("\nNo\tMasukan\t\tKeluaran")
	fmt.Printf("1\t%d\t\t%s\n", input, cetakDeret(input))
}
