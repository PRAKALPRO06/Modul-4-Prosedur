
package main

import "fmt"

func f1(a, b int) float64 {
    result := 2 * float64(a) - 0.5 * float64(b) + 0.3
    return result
}
func f2(a, b int, c *float64) {
    *c = 2 * float64(a) - 0.5 * float64(b) + 0.3
}

func main() {
    var a, b int
    var c float64
    fmt.Print("Enter two integers: ")
    fmt.Scan(&a, &b)
    f2(a, b, &c)
    fmt.Println("Result from f2 (stored in c): ", c)
    resultF1 := f1(b, a)
    fmt.Println("Result from f1:", resultF1)
}
