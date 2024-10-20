package main

import "fmt"

func sendEmailNotification(email string) {
	fmt.Println("mengirim email ke : Pendaftaran berhasil", email)
}

func main() {
	emails := []string{"Dimas@gmail.com", "Akal@gmail.com", "Hernanda@gmail.com"}
	fmt.Println("Mengirim email ke pengguna yang baru terdaftar: ")
	for _, email := range emails {
		sendEmailNotification(email)
	}
}
