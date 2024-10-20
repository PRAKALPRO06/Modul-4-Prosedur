//WISNU RANANTA RADITYA PUTRA (2311102013) IF-11-06

package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan a, b, c, d (pisahkan dengan spasi): ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	if a >= c && b >= d {
		p1_2311102013 := permutasi(a, c)
		c1_2311102013 := kombinasi(a, c)

		p2_2311102013 := permutasi(b, d)
		c2_2311102013 := kombinasi(b, d)

		fmt.Printf("P(a, c) = %d, K(a, c) = %d\n", p1_2311102013, c1_2311102013)
		fmt.Printf("P(b, d) = %d, K(b, d) = %d\n", p2_2311102013, c2_2311102013)
	} else {
		fmt.Println("Syarat a >= c dan b >= d harus dipenuhi.")
	}
}
