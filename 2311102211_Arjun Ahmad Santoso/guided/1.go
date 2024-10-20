package main

import "fmt"

func main() {
   var flag int
   var m  string	     
   fmt.Scan(&flag)
   fmt.Scan(&m)
   cetakPesan(m, flag)
}

func cetakPesan(m string, flag int) {
   var jenis string = "";
   if flag == 0 {
      jenis = "error"   
   } else if flag == 1 {
      jenis = "warning"
   } else if flag == 2 {
      jenis = "information"
   }
   fmt.Println(m, jenis);
}
