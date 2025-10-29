package main

import (
	"fmt"
	"strings"
)

func main() {
	var theaterName string = "Canda dan Tawa"
	const theaterTickets int = 50
	var remainingTickets uint = 50
	var bookingList = []string{}

	fmt.Printf("theaterTickets is %T, remainingTickets is %T, theaterName is %T\n", theaterTickets, remainingTickets, theaterName)

	fmt.Printf("Welcome to %v your theater booking system!\n", theaterName)
	fmt.Printf("We have a total of %v tickets in total, and %v are still available.\n", theaterTickets, remainingTickets)
	fmt.Println("Buy your tickets now!")

	for {

		if remainingTickets == 0 {
			fmt.Println("Sorry, all tickets are sold out for now!")
			fmt.Printf("All bookings: %v\n", bookingList)
			break
		}

		// user input part

		var userName string
		var userTickets int
		var email string
		var ticketType int

		fmt.Println("Please enter your name:")
		fmt.Scan(&userName)
		fmt.Println("Please enter your email address:")
		fmt.Scan(&email)
		fmt.Println("Hello", userName, ", how many tickets would you like to book?")
		fmt.Scan(&userTickets)

		// validation

		var isValidName bool = len(userName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNumber bool = userTickets > 0 && uint(userTickets) <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - uint(userTickets)
			bookingList = append(bookingList, userName)

			fmt.Println("Choose your ticket type:")
			fmt.Println("1. Regular")
			fmt.Println("2. VIP")
			fmt.Scan(&ticketType)

			var sectionName string
			switch ticketType {
			case 1:
				sectionName = "Regular Section"
			case 2:
				sectionName = "VIP Section"
			default:
				sectionName = "Regular Section (Default)"
				fmt.Println("Invalid choice, defaulting to Regular Section.")
			}

			fmt.Printf("Thank you %v for booking %v %v tickets. We will send the confirmation email at %v. Enjoy the show in %v!\n", userName, userTickets, sectionName, email, theaterName)
			fmt.Printf("%v tickets are still available for %v\n", remainingTickets, theaterName)

			fmt.Printf("The booking list is %v\n", bookingList)

		} else {
			fmt.Println("ERROR!")
			if !isValidName {
				fmt.Println("The name you entered is too short. Please try again.")
			} else if !isValidEmail {
				fmt.Println("The email address you entered is invalid. Please try again.")
			} else if !isValidTicketNumber {
				if userTickets <= 0 {
					fmt.Println("Please buy at least 1 ticket.")
				} else {
					fmt.Println("The number of tickets you requested exceeds the available tickets. Please try again.")
					fmt.Printf("Only %v tickets are remaining.\n", remainingTickets)
				}
			}
		}
	}
}
