package main

import "fmt"

func cetakDeret(n_226 int) {
   fmt.Printf("%d", n_226)
   
   for n_226 != 1 {
       if n_226 % 2 == 0 {
           n_226 = n_226 / 2
       } else {
           n_226 = 3 * n_226 + 1  
       }
       fmt.Printf(" %d", n_226)
   }
   fmt.Println()
}

func main() {
   var n int
   fmt.Printf("Masukkan nilai suku awal (n): ")
   fmt.Scan(&n)
   if n > 0 && n < 1000000{
	cetakDeret(n)
   }else {
	fmt.Println("Input tidak valid. Silakan masukkan bilangan bulat positif lebih kecil dari 1000000.")
   }
}
