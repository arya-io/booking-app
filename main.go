// Package main defines the entry point of the Go Conference booking application.
package main

import (
	"booking-app/helper" // Import helper package for validation functions
	"fmt"                // For formatted I/O
	"sync"               // For goroutines synchronization
	"time"               // For simulating email sending delay
)

// Conference details
var conferenceName = "Go Conference" // Conference name
const conferenceTickets int = 50     // Total tickets available
var remainingTickets uint = 50       // Tickets left for booking
var bookings = make([]UserData, 0)   // Slice to store bookings

// UserData stores attendee details
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // WaitGroup to handle goroutines

func main() {

	// Display welcome message and conference details
	greetUsers()

	// Get user input (name, email, number of tickets)
	firstName, lastName, email, userTickets := getUserInput()

	// Validate user input using helper package
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// Proceed with booking if inputs are valid
	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)    // Book the tickets
		wg.Add(1)                                              // Add a goroutine to WaitGroup
		go sendTicket(userTickets, firstName, lastName, email) // Send ticket asynchronously

		// Display the first names of current bookings
		firstNames := getFirstNames()
		fmt.Printf("These are the first names of all bookings: %v\n", firstNames)

		// Check if tickets are sold out
		if remainingTickets == 0 {
			fmt.Println("Our Conference is fully booked. Come back next year.")
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

	wg.Wait() // Wait for all goroutines to finish
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
		firstNames = append(firstNames, booking.firstName) // Collect first names from each booking
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
	remainingTickets -= userTickets // Deduct booked tickets from remaining tickets

	// Create user data and add it to bookings
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // Append user data to the bookings slice

	fmt.Printf("List of bookings: %v\n", bookings) // Print updated booking list

	// Confirm booking to the user
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

// sendTicket simulates sending a ticket via email
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) // Simulate a delay in sending the ticket
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("------------------------------------")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("------------------------------------")
	wg.Done() // Mark the goroutine as done
}
