//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import (
	"fmt"
)


func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}


func permutation(n, r int, hasil *int) {
	var nFact, nMinRFact int
	factorial(n, &nFact)
	factorial(n-r, &nMinRFact)
	*hasil = nFact / nMinRFact
}


func combination(n, r int, hasil *int) {
	var nFact, rFact, nMinRFact int
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &nMinRFact)
	*hasil = nFact / (rFact * nMinRFact)
}

func main() {

	var a, b, c, d int
	fmt.Print("Enter values for a, b, c, d: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)


	var P_ac, C_ac, P_bd, C_bd int


	permutation(a, c, &P_ac)
	combination(a, c, &C_ac)
	fmt.Printf("%d %d\n", P_ac, C_ac)


	permutation(b, d, &P_bd)
	combination(b, d, &C_bd)
	fmt.Printf("%d %d\n", P_bd, C_bd)
}
