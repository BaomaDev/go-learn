package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const confereceTickets = 50
	var remainingTickets = 50

	fmt.Println(" 	      _                           ")
	fmt.Println("	     | |                          ")
	fmt.Println("__      _____| | ___ ___  _ __ ___   ___ ")
	fmt.Println("\\ \\ /\\ / / _ \\ |/ __/ _ \\| '_ ` _ \\ / _ \\")
	fmt.Println(" \\ V  V /  __/ | (_| (_) | | | | | |  __/")
	fmt.Println("  \\_/\\_/ \\___|_|\\___\\___/|_| |_| |_|\\___|")
	fmt.Println("                                          ")
	fmt.Println("                                          ")
	fmt.Println("=========================================")
	fmt.Println("Get your Ticket to Attend", conferenceName)
	fmt.Println("=========================================")
	fmt.Println("Total Tickets:", confereceTickets, "Total Tickets Available:", remainingTickets)

}
