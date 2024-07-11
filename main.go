package main

import (
	"fmt"
	"go-tut/helper"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings []UserData

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for {

	firstName, lastName, email, userTickets := getUserInputs()

	isValidName, isValidEmail, isValidTickets := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		bookTickets(remainingTickets, firstName, lastName, email, userTickets)
		wg.Add(1)
		go sendTicketConfirmation(userTickets, email)

		firstNames := getFirstNames()

		fmt.Printf("The current bookings are:%v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("We are sold out! Thank you for your interest in the conference.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Printf("Invalid name. Please enter a valid first and last name.\n")
		}
		if !isValidEmail {
			fmt.Printf("Invalid email. Please enter a valid email.\n")
		}
		if !isValidTickets {
			fmt.Printf("Invalid number of tickets. Please enter a valid number of tickets.\n")
		}

		fmt.Printf("Invalid input. Please try again.\n")
	}
	wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("we have a total of %v tickets available. Hurry up and get yours! from the remaining %v tickets!\n", conferenceTickets, remainingTickets)
	fmt.Println(("Get your tickets here to attend!"))
}

func getUserInputs() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan((&lastName))

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan((&userTickets))

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets int, firstName string, lastName string, email string, userTickets int) {
	remainingTickets -= userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("current bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v! You have successfully purchased %v tickets to the %v conference. Please check your email %v for your confirmation.\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func sendTicketConfirmation(userTickets int, email string) {
	// code to send ticket confirmation
	time.Sleep((10 * time.Second))
	fmt.Println("############################################")
	fmt.Println("Sending ticket confirmation to user...")
	fmt.Printf("Ticket confirmation for %v tickets is sent to %v!\n", userTickets, email)
	fmt.Println("############################################")
	wg.Done()
}
