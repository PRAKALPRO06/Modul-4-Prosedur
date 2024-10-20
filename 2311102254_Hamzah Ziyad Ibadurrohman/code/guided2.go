package main

import "fmt"

func sendEmailNotification(email string) {
	fmt.Printf("Mengirim email ke %s: Pendaftaran berhasil.\n", email)
}

func main() {
	//daftar email pengguna baru
	emails := []string{"user1@example.com", "user2@example.com", "user3@example.com"}
	// Mengirim email ke pemberitahuan ke setiap pengguna
	fmt.Println("Mengirim email ke pengguna yang baru terdaftar:")
	for _, email := range emails {
		sendEmailNotification(email)
	}
}
