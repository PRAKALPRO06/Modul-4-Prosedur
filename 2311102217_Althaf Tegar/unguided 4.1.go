package main

import "fmt"

// Menghitung faktorial
func factorial(n_217 int) int {
	if n_217 == 0 {
		return 1
	}
	return n_217 * factorial(n_217-1)
}

// Menghitung permutasi
func permutation(n_217, r_217 int) int {
	return factorial(n_217) / factorial(n_217-r_217)
}

// Menghitung kombinasi
func combination(n_217, r_217 int) int {
	return factorial(n_217) / (factorial(r_217) * factorial(n_217-r_217))
}

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Permutasi dan kombinasi untuk a terhadap c
	p_a_c := permutation(a, c)
	c_a_c := combination(a, c)
	fmt.Printf("P(%d, %d) = %d\nC(%d, %d) = %d\n", a, c, p_a_c, a, c, c_a_c)

	// Permutasi dan kombinasi untuk b terhadap d
	p_b_d := permutation(b, d)
	c_b_d := combination(b, d)
	fmt.Printf("P(%d, %d) = %d\nC(%d, %d) = %d\n", b, d, p_b_d, b, d, c_b_d)
}
