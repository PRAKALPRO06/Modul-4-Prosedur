package main

import (
	"fmt"
)

func sendEmailNotification(email string) {
	fmt.Printf("Mengirim email ke: %s\n", email)
}

func main() {
	emails := []string{"User1@example.com", "User2@example.com", "User3@example.com"}

	fmt.Println("Mengirim Email ke yang baru terdaftar:")

	for _, email := range emails {
		sendEmailNotification(email)
	}
}
