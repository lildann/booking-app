package main

import (
	"fmt"
	"strings"
)

// Package level variables, accessible to all functions in this package
const conferenceTickets = 50

var conferenceName = "~ Go Conference ~"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := returnFirstNamesOnly()
			fmt.Printf("These are the first names of the bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end progam
				fmt.Println("No more tickets remain, come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email must contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("There are %v tickets remaining; you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) // Scan() used for i/o functionality to get user input

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func returnFirstNamesOnly() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Hello, %s %s! You have booked %d tickets to the %s.\n", firstName, lastName, userTickets, conferenceName)
	fmt.Printf("You will receive a booking confirmation at %s\n", email)
	fmt.Printf("%v tickets are still available\n", remainingTickets)
}
