package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// Package level variables, accessible to all functions in this package
const conferenceTickets = 50

var conferenceName = "~ Go Conference ~"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
// Turn bookings slice from a list of strings to a list of Structs

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			// go routine creates concurrency

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
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Hello, %s %s! You have booked %d tickets to the %s.\n", firstName, lastName, userTickets, conferenceName)
	fmt.Printf("You will receive a booking confirmation at %s\n", email)
	fmt.Printf("%v tickets are still available\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // emulating delay, stops the execution of the thread

	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n%v \nto email address: %v\n", ticket, email)
	fmt.Println("###########")
}
