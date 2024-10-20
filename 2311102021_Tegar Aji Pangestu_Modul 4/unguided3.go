package main 
import "fmt"

//Tegar Aji Pangestu
func cetakDeret(n int) {
	fmt.Print(n, " ")
	prev := n
	current := n
	for i:=0; i<n; i++ {
		if prev % 2 == 0 {
			current = prev / 2
			
		} else {
			current = 3*prev + 1
		}
		prev = current
		fmt.Print(current, " ")
		if current == 1 { break }
	}
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	if n > 1000000 {
		fmt.Println("Input tidak boleh melebihi 1000000")
		return
	}
	cetakDeret(n)
}