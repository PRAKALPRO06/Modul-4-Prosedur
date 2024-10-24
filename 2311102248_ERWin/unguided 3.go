package main

import "fmt"

func deretBilangan(n int) {
	win2248 := n

	for win2248 != 1 {
		fmt.Print(win2248, " ")

		if win2248%2 == 0 {
			win2248 = win2248 / 2
		} else {
			win2248 = (3 * win2248) + 1
		}
	}
	fmt.Println(win2248)
}

func main() {
	var n int

	fmt.Scan(&n)

	deretBilangan(n)
}
