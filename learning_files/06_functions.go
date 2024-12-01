// Package main defines the entry point of the Go Conference booking application.
package main

import (
	"fmt"     // Provides formatted I/O functions
	"strings" // For string operations (used to split names)
)

func main() {
	// Conference details
	conferenceName := "Go Conference" // Conference name
	const conferenceTickets int = 50  // Total number of tickets
	var remainingTickets uint = 50    // Tickets currently available
	bookings := []string{}            // Stores booked first and last names

	// Display welcome message
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// Infinite loop for handling bookings
	for {

		// Get user input
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Process booking if input is valid
		if isValidName && isValidEmail && isValidTicketNumber {

			// Book tickets
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// Display all first names of current bookings
			firstNames := getFirstNames(bookings)
			fmt.Printf("These are the first names of all bookings: %v\n", firstNames)

			// Stop when tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our Conference is fully booked. Come back next year.")
				break
			}
		} else {
			// Handle invalid input cases
			if !isValidName {
				fmt.Println("The first or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address is invalid (missing '@').")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets entered is invalid.\n")
			}
		}
	}
}

// greetUsers displays the initial welcome message and conference info
func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets now to secure your spot!")
}

// getFirstNames extracts and returns the first names of all booked attendees
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)          // Split full name into fields
		firstNames = append(firstNames, names[0]) // Append first name to the list
	}
	return firstNames
}

// validateUserInput checks if user input (name, email, ticket count) is valid
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2      // Validate name length
	isValidEmail := strings.Contains(email, "@")                  // Validate email format
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // Validate ticket count
	return isValidName, isValidEmail, isValidTicketNumber
}

// getUserInput prompts the user for personal details and ticket count
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	// Collect user details
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// bookTicket handles ticket booking, updates the booking list, and prints confirmation
func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets -= userTickets                     // Deduct booked tickets
	bookings = append(bookings, firstName+" "+lastName) // Add booking to the list

	// Print booking confirmation
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
