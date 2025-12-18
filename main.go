package main

import (
	"fmt"
	"strings"
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
func greet() {
	fmt.Println("Welcome to the", confname, "booking appli cation")
	fmt.Println("We have a total of", conftickets, "tickets and", remainingtickets, "are still available.")
	fmt.Println("Get your tickets to attend")
}
func getsFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func validateUserInput(firstname string, lastname string, email string, usertickets int, remainingtickets int) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := usertickets > 0 && usertickets <= remainingtickets
	return isValidName, isValidEmail, isValidTicketNumber
}
func bookTicket(usertickets int, firstname string, lastname string, email string) {
	remainingtickets = remainingtickets - usertickets
	bookings = append(bookings, firstname+" "+lastname)
	fmt.Printf("Thank you %v %v for booking %v tickets.\n You will receive a confirmation email at %v\n", firstname, lastname, usertickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingtickets, confname)
}

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
