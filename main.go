package main

import "fmt"

func main() {
	var theaterName string = "Canda dan Tawa"
	const theaterTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("theaterTickets is %T, remainingTickets is %T, theaterName is %T\n", theaterTickets, remainingTickets, theaterName)

	fmt.Printf("Welcome to %v your theater booking system!\n", theaterName)
	fmt.Printf("We have a total of %v tickets in total, and %v are still available.\n", theaterTickets, remainingTickets)
	fmt.Println("Buy your tickets now!")

	// user input part

	var userName string
	var userTickets int
	var email string

	fmt.Println("Please enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Hello", userName, ", how many tickets would you like to book?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v for booking %v tickets. We will send the confirmation email at %v. Enjoy the show in %v!\n", userName, userTickets, email, theaterName)
	fmt.Printf("%v tickets are still available for %v\n", remainingTickets, theaterName)

}
