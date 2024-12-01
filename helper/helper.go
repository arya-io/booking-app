package helper

import "strings"

// ValidateUserInput checks if the user's name, email, and ticket count are valid
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2                  // Names must be at least 2 characters
	isValidEmail := strings.Contains(email, "@")                              // Email must include '@'
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // Ensure valid ticket number
	return isValidName, isValidEmail, isValidTicketNumber
}
