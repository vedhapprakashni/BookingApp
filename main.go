package main

import (
	"fmt"
)

func main() {
	var confname = "Go conference"
	const conftickets = 50
	var remainingtickets = 50

	fmt.Printf("conferencetickets is %T,remaining tickets is %T,confname is %T\n", conftickets, remainingtickets, confname)
	// This is a placeholder for the main function.

	fmt.Println("welcome to the", confname, "booking application")
	fmt.Println("We have a total of", conftickets, "tickets and", remainingtickets, "are still available.")
	fmt.Println("Get your tickets to attend")

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

	fmt.Println("Enter your userticket:")
	fmt.Scan(&usertickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstname, lastname, usertickets, email)

}
