# Booking App

This is a simple Go-based application to manage ticket bookings for the "Go Conference." The app allows users to book conference tickets by providing their personal information and number of tickets, validates the input, and sends a confirmation email (simulated with a delay).

## Features

- Validate user input (name, email, ticket count)
- Book conference tickets
- Show remaining tickets
- Display the list of first names of all booked attendees
- Asynchronous ticket confirmation email simulation

## Prerequisites

Before running the application, make sure you have the following installed on your machine:

- **Go (Golang)**
- A working terminal or command prompt.

You can verify if Go is installed correctly by running the following command in your terminal:

```bash
go version
```

## Installation

1. **Clone the repository** to your local machine:

```bash
git clone https://github.com/arya-io/booking-app.git
```
2. **Navigate to the project directory:**

```bash
cd booking-app
```

## Usage

To run the application, execute the following command in the project directory:

```bash
go run main.go
```

## Interaction
Once the app is running, you'll be prompted to provide the following information:

1. **First Name**: Your first name.
2. **Last Name**: Your last name.
3. **Email**: Your email address (must contain `@`).
4. **Number of Tickets**: The number of tickets you wish to book.
The app will validate the input, book the tickets if valid, and simulate sending a confirmation email after a short delay.

## Example

```bash
Enter your first name: John
Enter your last name: Doe
Enter your email address: john.doe@example.com
Enter the number of tickets: 2
```

## Project Structure

```bash

booking-app/ │ 
             ├── main.go # The main entry point for the application 
             ├── helper.go # Contains helper functions for input validation 
             ├── go.mod # Go module file for dependency management 
             ├── go.sum # Checksums for module dependencies 
             └── README.md # Project documentation (this file)
```


- **main.go**: Handles the core logic of the application, including booking tickets, validating input, and sending confirmation emails.
- **helper.go**: Contains reusable functions to validate user input (name, email, and ticket count).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.