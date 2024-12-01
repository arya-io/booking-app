// Package main defines the entry point of the Go Conference booking application.
package main

import (
	"fmt"     // For formatted input/output
	"strings" // For string manipulation
)

func main() {
	// Conference details
	conferenceName := "Go Conference" // Name of the conference
	const conferenceTickets int = 50  // Total available tickets
	var remainingTickets uint = 50    // Tickets still available
	bookings := []string{}            // List of bookings (first + last names)

	// Welcome message
	fmt.Printf("Welcome to %v Booking Application!\n", conferenceName)
	fmt.Printf("We have %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now to secure your spot!")

	// Infinite loop to take bookings until tickets run out
	for {
		// User details
		var firstName, lastName, email string
		var userTickets uint

		// Get user input
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		// Validate input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// If valid, process the booking
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets                     // Update remaining tickets
			bookings = append(bookings, firstName+" "+lastName) // Add booking

			// Confirmation message
			fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n",
				firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			// Display first names of all attendees
			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)          // Split full name
				firstNames = append(firstNames, names[0]) // Collect first name
			}
			fmt.Printf("These are the first names of all bookings: %v\n", firstNames)

			// Stop if no tickets left
			if remainingTickets == 0 {
				fmt.Println("Our Conference is fully booked. Come back next year.")
				break
			}
		} else {
			// Handle invalid inputs
			if !isValidName {
				fmt.Println("First or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address is invalid (missing @ symbol).")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets is invalid.\n")
			}
		}
	}
}
