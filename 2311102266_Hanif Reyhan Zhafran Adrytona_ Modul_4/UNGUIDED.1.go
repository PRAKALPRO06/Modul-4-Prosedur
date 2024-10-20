package main

import (
	"fmt"
)

func main() {
	var x, y, z, w int
	fmt.Print("Masukkan x, y, z, w: ")
	fmt.Scanf("%d %d %d %d", &x, &y, &z, &w)

	if x >= z && y >= w {

		permutasi1_2311102266 := permutasi(x, z)
		kombinasi1_2311102266 := kombinasi(x, z)
		permutasi2_2311102266 := permutasi(y, w)
		kombinasi2_2311102266 := kombinasi(y, w)

		fmt.Printf("Permutasi P(x, z) = %d, Kombinasi C(x, z) = %d\n", permutasi1_2311102266, kombinasi1_2311102266)
		fmt.Printf("Permutasi P(y, w) = %d, Kombinasi C(y, w) = %d\n", permutasi2_2311102266, kombinasi2_2311102266)
	} else {
		fmt.Println("Syarat x >= z dan y >= w harus dipenuhi.")
	}
}
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}
func permutasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

// 2311102266 - Hanip 11-if-06
