package main

import "strings"

// Fungsi untuk memvalidasi input pengguna
func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	// Memastikan nama depan dan belakang minimal 2 karakter
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// Memastikan email mengandung simbol '@'
	isValidEmail := strings.Contains(email, "@")

	// Memastikan jumlah tiket yang diminta tidak melebihi tiket yang tersedia
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
