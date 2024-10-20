//2311102018 Aryo Tegar Sukarno
package main

import (
        "fmt"
)

func factorial(n int) int {
        if n == 0 {
                return 1
        }
        return n * factorial(n-1)
}

func permutation(n, r int) int {
        return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
        return permutation(n, r) / factorial(r)
}

func main() {
        var a, b, c, d int
        fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

        
        p1 := permutation(a, c)
        c1 := combination(a, c)
      
        p2 := permutation(b, d)
        c2 := combination(b, d)
        
        fmt.Println(p1, c1)
        fmt.Println(p2, c2)
}