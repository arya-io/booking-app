// Package main defines the entry point of the Go Conference booking application.
package main

import (
	"fmt"     // For formatted I/O
	"strings" // For string operations like splitting names
)

// Conference details
var conferenceName = "Go Conference" // Conference name
const conferenceTickets int = 50     // Total tickets available
var remainingTickets uint = 50       // Tickets left for booking
var bookings = []string{}            // List to store booked attendees' names

func main() {

	// Display welcome message and details
	greetUsers()

	// Run the booking loop until tickets sell out
	for {
		// Collect user input (name, email, number of tickets)
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// If all inputs are valid, proceed with booking
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email) // Book tickets

			// Show the first names of all bookings
			firstNames := getFirstNames()
			fmt.Printf("These are the first names of all bookings: %v\n", firstNames)

			// End the program if no tickets are left
			if remainingTickets == 0 {
				fmt.Println("Our Conference is fully booked. Come back next year.")
				break
			}
		} else {
			// Handle invalid inputs by displaying relevant messages
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

// greetUsers prints a welcome message with conference details
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now to secure your spot!")
}

// getFirstNames returns a list of first names of all bookings
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)          // Split full name
		firstNames = append(firstNames, names[0]) // Store first name
	}
	return firstNames
}

// validateUserInput checks if the user input is valid
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2                  // Name must have at least 2 characters
	isValidEmail := strings.Contains(email, "@")                              // Email must contain '@'
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // Ticket count must be valid
	return isValidName, isValidEmail, isValidTicketNumber
}

// getUserInput prompts the user for name, email, and ticket count
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	// Prompt for user details
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

// bookTicket deducts tickets and confirms the booking
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets                     // Update available tickets
	bookings = append(bookings, firstName+" "+lastName) // Add to bookings

	// Print booking confirmation
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
