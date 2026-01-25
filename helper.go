package main

import (
	"fmt"
	"strings"
)

// greet() prints the welcome message and current ticket status
// It uses global variables like confname, conftickets, remainingtickets
func greet() {
	fmt.Println("Welcome to the", confname, "booking application")
	fmt.Println("We have a total of", conftickets, "tickets and", remainingtickets, "are still available.")
	fmt.Println("Get your tickets to attend")
}

// getFirstNames() returns a slice containing only the first names
// from the bookings slice
func getsFirstNames() []string {
	firstNames := []string{} // empty slice to store first names

	// Loop through all bookings
	for _, booking := range bookings {
		// Split "FirstName LastName" into words
		var names = strings.Fields(booking)

		// Add only the first name to the slice
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

// validateUserInput() checks whether user input is valid
// Returns three booleans:
// 1. isValidName → checks name length
// 2. isValidEmail → checks if email contains "@"
// 3. isValidTicketNumber → checks ticket availability
func validateUserInput(
	firstname string,
	lastname string,
	email string,
	usertickets int,
	remainingtickets int,
) (bool, bool, bool) {

	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := usertickets > 0 && usertickets <= remainingtickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// bookTicket() books the tickets for the user
// - reduces remaining tickets
// - stores booking details
// - prints confirmation message
func bookTicket(usertickets int, firstname string, lastname string, email string) {

	// Reduce available tickets
	remainingtickets = remainingtickets - usertickets

	// Save booking as "FirstName LastName"
	bookings = append(bookings, firstname+" "+lastname)

	// Confirmation message
	fmt.Printf(
		"Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n",
		firstname, lastname, usertickets, email,
	)

	fmt.Printf("%v tickets remaining for %v\n", remainingtickets, confname)
}

// getUserInput() takes input from the user.
// Returns first name, last name, email, and ticket count.
func getuserInput() (string, string, string, int) {

	var firstname string
	var lastname string
	var email string
	var usertickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&usertickets)

	return firstname, lastname, email, usertickets
}
