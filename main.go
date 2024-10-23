package main

import "fmt"

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"

	fmt.Println(" 	      _                           ")
	fmt.Println("	     | |                          ")
	fmt.Println("__      _____| | ___ ___  _ __ ___   ___ ")
	fmt.Println("\\ \\ /\\ / / _ \\ |/ __/ _ \\| '_ ` _ \\ / _ \\")
	fmt.Println(" \\ V  V /  __/ | (_| (_) | | | | | |  __/")
	fmt.Println("  \\_/\\_/ \\___|_|\\___\\___/|_| |_| |_|\\___|")
	fmt.Println("                                          ")
	fmt.Println("                                          ")
	fmt.Println("=========================================")
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("=========================================")
	fmt.Printf("We have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceTickets, remainingTickets)

	//User Inputs
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	// book ticket in system
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
