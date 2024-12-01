// Package main defines the entry point of the Go Conference booking application.
package main

import (
	"fmt"     // For formatted I/O
	"strings" // For string manipulation
)

func main() {
	// Conference details
	conferenceName := "Go Conference"   // Name of the conference
	const conferenceTickets int = 50    // Total number of tickets
	var remainingTickets uint = 50      // Tickets left
	bookings := []string{}              // Store bookings as a list of names

	// Display welcome message
	fmt.Printf("Welcome to %v Booking Application!\n", conferenceName)
	fmt.Printf("We have %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now to secure your spot!")

	// Infinite loop to continuously take bookings until tickets run out
	for {
		var firstName, lastName, email string // User details
		var userTickets uint                  // Tickets to be booked by the user

		// Collect user details
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		// Check if the requested number of tickets is available
		if userTickets <= remainingTickets {
			// Process booking
			remainingTickets -= userTickets                     // Deduct booked tickets from remaining
			bookings = append(bookings, firstName+" "+lastName) // Save full name to bookings list

			// Confirm booking and remaining tickets
			fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n",
				firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			// Extract and display first names of all attendees
			firstNames := []string{} // Store first names
			for _, booking := range bookings {
				names := strings.Fields(booking)          // Split full name into parts
				firstNames = append(firstNames, names[0]) // Append first name to the list
			}
			fmt.Printf("These are the first names of all bookings: %v\n", firstNames)

			// Stop taking bookings when all tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
			}
		} else {
			// Handle case where requested tickets exceed remaining
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n",
				remainingTickets, userTickets)
		}
	}
}
