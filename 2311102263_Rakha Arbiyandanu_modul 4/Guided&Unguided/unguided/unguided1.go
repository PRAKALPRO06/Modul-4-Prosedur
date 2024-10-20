package main

import (
	"fmt"
)


func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}


func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}


func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Println("a >= c dan b >= d")


	fmt.Print("Masukkan empat bilangan bulat (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)


	if a < c || b < d {
		fmt.Println("Syarat a >= c dan b >= d tidak terpenuhi.")
		return
	}


	p1 := permutation(a, c)
	c1 := combination(a, c)


	p2 := permutation(b, d)
	c2 := combination(b, d)


	fmt.Println("Hasil untuk (a, c):")
	fmt.Printf("Permutasi P(%d, %d) = %d\n", a, c, p1)
	fmt.Printf("Kombinasi C(%d, %d) = %d\n", a, c, c1)
 
	fmt.Println("Hasil untuk (b, d):")
	fmt.Printf("Permutasi P(%d, %d) = %d\n", b, ,d p2)
	fmt.Printf("Kombinasi C(%d, %d) = %d\n", b, d, c2)
}
