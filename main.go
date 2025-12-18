package main

import (
	"fmt"
)

var confname = "Go conference"

const conftickets = 50

var remainingtickets = 50
var bookings []string

func main() {

	greet()

	for {
		firstname, lastname, email, usertickets := getuserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstname, lastname, email, usertickets, remainingtickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(usertickets, firstname, lastname, email)

			// Extract first names
			firstNames := getsFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingtickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("First name or last name you entered is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address you entered doesn't contain @ sign\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid\n")
			}

		}

	}
}
