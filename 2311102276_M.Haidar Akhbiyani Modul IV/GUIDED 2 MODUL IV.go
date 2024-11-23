package main

import "fmt"

func sendEmailNotification(email string) {
	fmt.Printf("mengirim email ke %s: Pendaftaran berhasil.\n" , email)
}

func main() {
	emails := []string{"Haidar@example.com", "Akhbi@example.com", "Yani@example.com"}
	fmt.Println("mengirim email ke pengguna yang baru daftar:")
	for _, email := range emails {
		sendEmailNotification(email)
	}
}