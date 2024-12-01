// Package main defines the entry point of the Go Conference booking application.
package main

import (
	"booking-app/helper" // Import helper package for validation functions
	"fmt"                // For formatted I/O
	"strconv"            // For string conversion
)

// Conference details
var conferenceName = "Go Conference"        // Conference name
const conferenceTickets int = 50            // Total tickets available
var remainingTickets uint = 50              // Tickets left for booking
var bookings = make([]map[string]string, 0) // List to store booking details (maps for each user)

func main() {

	// Display welcome message and conference details
	greetUsers()

	// Loop to handle bookings until tickets are sold out
	for {
		// Get user input (name, email, number of tickets)
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user input using helper package
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Proceed with booking if all inputs are valid
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email) // Book the tickets

			// Display all first names of current bookings
			firstNames := getFirstNames()
			fmt.Printf("These are the first names of all bookings: %v\n", firstNames)

			// If tickets are sold out, display a message and exit the loop
			if remainingTickets == 0 {
				fmt.Println("Our Conference is fully booked. Come back next year.")
				break
			}
		} else {
			// Handle invalid inputs by showing appropriate messages
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
		firstNames = append(firstNames, booking["firstName"]) // Collect first names from each booking
	}
	return firstNames
}

// getUserInput prompts the user for name, email, and number of tickets
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

// bookTicket deducts tickets, updates booking list, and confirms booking
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets // Deduct the booked tickets from remaining tickets

	// Create a map to store user booking details
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10) // Convert ticket number to string

	bookings = append(bookings, userData) // Add user data to the list of bookings

	fmt.Printf("List of bookings: %v\n", bookings) // Print updated booking list

	// Confirm booking to the user
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
