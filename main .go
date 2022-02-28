<<<<<<< Updated upstream
package main

import (
	"fmt"
	"strings"
	"time"
)

const conferenceTicket int = 50

var conferenceName = "Go Conference"
var remainingTicket uint = 50
var bookigns = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := validUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {

			if userTickets < remainingTicket {
				bookTicket(userTickets, firstName, lastName, email)
				go sendTicket(userTickets, firstName, lastName, email)

				//get first Name
				firstName := printFirstName()

				fmt.Println("The first Names of Bookings are :\n", firstName)

				noTicketsAvailable := remainingTicket == 0

				if noTicketsAvailable {
					fmt.Println("Our Conference is booked our. Come back next Year")
					break

				}

			} else {
				fmt.Println("we only have", remainingTicket, "Tickets remaining, so you can't book", userTickets)

			}

		} else {
			if !isValidName {
				fmt.Println("Fist Name and Last Name must be atleast 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Email Address is not valid")

			}
			if !isValidTickets {
				fmt.Println("Number of Tickets You enter is invalid")
			}
			fmt.Println("Invalid Input Please Try Again")
		}

	}
}

func greetUsers() {

	fmt.Printf("Welcome to our %v \n", conferenceName)
	fmt.Println("We Have total of ", conferenceTicket, " Tickets and ", remainingTicket, " are still Available")
	fmt.Println("Get Your Tickets here to attend")
}

func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookigns {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name:")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last Name:")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your Email Address:")
	fmt.Scanln(&email)

	fmt.Println("Enter the Number of Tickets:")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookigns = append(bookigns, userData)
	fmt.Println("Complete List of Booking", bookigns)

	fmt.Printf("Thank You %v %v for Booking %v tickets. You Will recive  a confirmation email at  %v \n", firstName, lastName, userTickets, email)
	fmt.Println(remainingTicket, "Tickets remaining for ", conferenceName)

}

func validUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Ticket for %v %v ", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending Ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("#######################")
}
=======
>>>>>>> Stashed changes
