package main 

import "fmt"

func faktorial(a int) int {
	result := 1
	for i:=1; i<=a; i++ {
		result = result * i
	}
	return result
}
func permutasi(n, r int) int {
	return faktorial(n)/faktorial(n-r)
}
func kombinasi(n, r int) int {
	return faktorial(n)/(faktorial(n-r)*faktorial(r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a < c || b < d {
		fmt.Println("Input is not valid")
		return
	}
	fmt.Print("\n")
	fmt.Print(permutasi(a, c), kombinasi(a, c), "\n")
	fmt.Print(permutasi(b, d), kombinasi(b, d), "\n")
}