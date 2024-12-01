// Package main defines the entry point of the Go Conference booking application.
package main

import "fmt" // Import the fmt package for formatted I/O

// main is the entry point of the program.
func main() {
	// Conference details
	conferenceName := "Go Conference" // Name of the conference
	const conferenceTickets int = 50  // Total tickets (constant)
	var remainingTickets uint = 50    // Remaining tickets (variable)

	// Welcome message
	fmt.Printf("Welcome to %v Booking Application!\n", conferenceName)
	fmt.Printf("We have %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now to secure your spot!")

	// User input variables
	var firstName string // User's first name
	var lastName string  // User's last name
	var email string     // User's email address
	var userTickets uint // Number of tickets to book

	// Collect user data
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	// Update remaining tickets
	remainingTickets -= userTickets

	// Confirmation message
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
