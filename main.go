package main

import "fmt"

func main() {
	var conferenceName = "~ Go Conference ~"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// booking array with known array size of 50:
	bookings := []string{}

	fmt.Printf("Welcome to the %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets -= userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first element: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Hello, %s %s! You have booked %d tickets to the %s.\n", firstName, lastName, userTickets, conferenceName)
	fmt.Printf("You will receive a booking confirmation at %s\n", email)

	fmt.Printf("%v tickets are still available\n", remainingTickets)

}
