// Package main defines the entry point of the Go Conference booking application.
package main

import "fmt" // Import fmt for formatted I/O

// main is the starting point of the program.
func main() {
	// Conference details
	conferenceName := "Go Conference" // Conference name
	const conferenceTickets int = 50  // Total tickets (constant)
	var remainingTickets uint = 50    // Tickets available (variable)
	bookings := []string{}            // List of bookings

	// Welcome message
	fmt.Printf("Welcome to %v Booking Application!\n", conferenceName)
	fmt.Printf("We have %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now to secure your spot!")

	// User input variables
	var firstName string // First name input
	var lastName string  // Last name input
	var email string     // Email input
	var userTickets uint // Number of tickets requested

	// Collect user information
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName) // Read first name

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName) // Read last name

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email) // Read email

	fmt.Print("Enter the number of tickets: ")
	fmt.Scan(&userTickets) // Read tickets

	// Update ticket count and record booking
	remainingTickets -= userTickets                     // Deduct booked tickets
	bookings = append(bookings, firstName+" "+lastName) // Save booking

	// Booking confirmation
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// Show current bookings
	fmt.Printf("These are all our bookings: %v\n", bookings)
}
