package main

import (
	"fmt"
)

// Factorial function
func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Permutation function
func permutation(a, b int) int {
	kiri := factorial(a)
	kanan := factorial(a - b)
	return kiri / kanan
}

// Combination function
func combination(a, b int) int {
	kiri := factorial(a)
	kanan1 := factorial(b)
	kanan2 := factorial(a - b)
	return kiri / (kanan1 * kanan2)
}

// TaskDiskrit handles permutation and combination calculations
func taskDiskrit(a, b, c, d int) {
	fmt.Printf("Permutasian dari (%d,%d): %d\n", a, c, permutation(a, c))
	fmt.Printf("Kombinasi dari (%d, %d): %d\n", a, c, combination(a, c))
	fmt.Printf("Permutasian dari (%d,%d): %d\n", b, d, permutation(b, d))
	fmt.Printf("Kombinasi dari (%d, %d): %d\n", b, d, combination(b, d))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukan Nilai untuk a, b, c, dan d: ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		taskDiskrit(a, b, c, d)
	} else {
		fmt.Println("Invalid input: a harus >= c and b harus >= d.")
	}
}
