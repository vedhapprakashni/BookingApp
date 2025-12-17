package main

import (
	"fmt"
	"strings"
)

func main() {
	var confname = "Go conference"
	const conftickets = 50
	var remainingtickets = 50
	var bookings []string

	greet(confname, conftickets, remainingtickets)

	for {
		firstname, lastname, email, usertickets := getuserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstname, lastname, email, usertickets, remainingtickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			remainingtickets = remainingtickets - usertickets
			bookings = append(bookings, firstname+" "+lastname)

			fmt.Printf(
				"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
				firstname, lastname, usertickets, email,
			)

			fmt.Printf("%v tickets are remaining for %v\n", remainingtickets, confname)

			// Extract first names
			firstNames := getsFirstNames(bookings)
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
		// city := "London"

		// switch city {
		// case "New York":
		// 	fmt.Println("The city that never sleeps")
		// case "Singapore", "Hong Kong":
		// 	fmt.Println("One of the busiest ports in the world")
		// case "London", "Berlin":
		// 	fmt.Println("The capital of England")
		// case "Mexcio City":
		// 	fmt.Println("The largest city in the Americas")
		// default:
		// 	fmt.Println("No special information about this city")

	}
}
func greet(confname string, conftickets int, remainingtickets int) {
	fmt.Println("Welcome to the", confname, "booking appli cation")
	fmt.Println("We have a total of", conftickets, "tickets and", remainingtickets, "are still available.")
	fmt.Println("Get your tickets to attend")
}
func getsFirstNames(bookings []string) []string {
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
