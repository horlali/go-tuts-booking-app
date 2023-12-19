package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName string = "GoLang Conference"
const conferenceTickets int = 50

var remainingTickets uint = 50
var booking_list = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
	// isOptedForNewsletter bool
}

var wg = sync.WaitGroup{}

func main() {

	println("==================================")
	greetings()

	println("**********************************")
	fmt.Printf("The remaining tickets for this conference is %v\n", remainingTickets)

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(firstName, lastName, email, userTickets)

		attendeelist := getAttendeesName()
		fmt.Printf("Attendees of the conference are: %v\n \n", attendeelist)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}

	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email is invalid")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid number of tickets")
		}
	}

	wg.Wait()
}

func greetings() {
	fmt.Printf("Welcome to the %v ticket booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets for this conference\n", conferenceTickets)
	fmt.Printf("Get your tickets here to attend\n \n")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	booking_list = append(booking_list, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("We will send you a confirmation to your email: %v\n", email)
	fmt.Printf("Remaining tickets for the event is: %v\n", remainingTickets)
}

func getAttendeesName() []string {
	attendeesFirstNames := []string{}

	for _, booking := range booking_list {
		attendeesFirstNames = append(attendeesFirstNames, booking.firstName)
	}

	return attendeesFirstNames
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	println("//////////////////////////////////")
	fmt.Printf("Sending tickets:\n%v \nto email address %v\n", ticket, email)
	println("//////////////////////////////////")
	wg.Done()
}
