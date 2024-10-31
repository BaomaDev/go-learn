package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50 // Jumlah tiket yang tersedia

var conferenceName = "Go Conference" // Nama konferensi
var remainingTickets uint = 50       // Tiket yang masih tersedia
var bookings = make([]UserData, 0)   // Daftar pemesanan

// Struct untuk menyimpan data pemesanan pengguna
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// WaitGroup untuk mengelola goroutine
var wg = sync.WaitGroup{}

func main() {
	displayWelcomeMessage()

	// Loop untuk menerima pemesanan terus-menerus sampai tiket habis
	for {
		// Mengumpulkan input dari pengguna
		firstName, lastName, email, userTickets := getUserInput()

		// Validasi input menggunakan fungsi dari helper.go
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// Jika semua input valid, lanjutkan pemesanan
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email) // Memproses pemesanan tiket

			// Menambahkan goroutine untuk mengirim tiket secara asynchronous
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// Menampilkan nama depan dari semua yang telah memesan tiket
			firstNames := getFirstNames()
			fmt.Printf("Nama-nama depan dari pemesanan adalah: %v\n", firstNames)

			// Keluar jika tiket sudah habis
			if remainingTickets == 0 {
				fmt.Println("Konferensi sudah penuh. Silakan kembali tahun depan!")
				break
			}
		} else {
			// Menampilkan pesan error sesuai hasil validasi
			if !isValidName {
				fmt.Println("Nama depan atau belakang yang Anda masukkan terlalu pendek.")
			}
			if !isValidEmail {
				fmt.Println("Alamat email harus mengandung simbol '@'.")
			}
			if !isValidTicketNumber {
				fmt.Println("Jumlah tiket yang dimasukkan tidak valid.")
			}
		}
	}
	wg.Wait() // Menunggu semua goroutine selesai sebelum program berakhir
}

// Fungsi untuk menampilkan pesan sambutan dengan ASCII art
func displayWelcomeMessage() {
	fmt.Println(" 	      _                           ")
	fmt.Println("	     | |                          ")
	fmt.Println("__      _____| | ___ ___  _ __ ___   ___ ")
	fmt.Println("\\ \\ /\\ / / _ \\ |/ __/ _ \\| '_  _ \\ / _ \\")
	fmt.Println(" \\ V  V /  __/ | (_| (_) | | | | | |  __/")
	fmt.Println("  \\_/\\_/ \\___|_|\\___\\___/|_| |_| |_|\\___|")
	fmt.Println("                                          ")
	fmt.Println("                                          ")
	fmt.Println("=========================================")
	fmt.Printf("Selamat datang di aplikasi pemesanan %v.\n", conferenceName)
	fmt.Println("=========================================")
	fmt.Printf("Kami memiliki total %v tiket dan %v tiket masih tersedia.\nDapatkan tiket Anda di sini untuk hadir!\n", conferenceTickets, remainingTickets)
}

// Fungsi untuk mendapatkan nama depan dari semua pengguna yang telah memesan tiket
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// Fungsi untuk mengumpulkan input dari pengguna
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Masukkan Nama Depan Anda: ")
	fmt.Scanln(&firstName)

	fmt.Println("Masukkan Nama Belakang Anda: ")
	fmt.Scanln(&lastName)

	fmt.Println("Masukkan Alamat Email Anda: ")
	fmt.Scanln(&email)

	fmt.Println("Masukkan Jumlah Tiket: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

// Fungsi untuk memproses pemesanan tiket
func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets

	// Membuat entri pemesanan baru
	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Terima kasih %v %v telah memesan %v tiket. Anda akan menerima email konfirmasi di %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tiket tersisa untuk %v\n", remainingTickets, conferenceName)
}

// Goroutine untuk mensimulasikan pengiriman tiket
func sendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(5 * time.Second) // Simulasi penundaan dalam pengiriman tiket
	ticket := fmt.Sprintf("%v tiket untuk %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Mengirim tiket:\n%v\nke alamat email %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() // Menandai goroutine telah selesai
}
