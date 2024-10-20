package main

import "fmt"

func sendEmailNotification(email string) {
	fmt.Printf("mengirim email ke %s: Pendaftaran berhasil.\n" , email)
}

func main() {
	emails := []string{"user1@example.com", "user2@example.com", "user3@example.com"}
	fmt.Println("mengirim email ke pengguna yang baru daftar:")
	for _, email := range emails {
		sendEmailNotification(email)
	}
}