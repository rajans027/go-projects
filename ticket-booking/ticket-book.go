package main

import (
	"fmt"
	"strconv"
)

var toolName string = "Smooth"

const totalTickets int = 80

var remainingTickets int = 80

//user provided info stored as variables

var userName string
var firstName string
var lastName string
var userTickets int

//var bookingsName []string //definfing array variables and its datatype

var bookingsName = make([]map[string]string, 0) //creating an empty list of map
var nickNames []string

func main() {

	welcomeMessage(toolName, totalTickets, remainingTickets)

	for {

		// function to gather user input
		firstName, lastName, userTickets := gatherUserInput() //all three values are return values from the function and stored here to be used in main faction

		//function to check the validity of the user input, return values are used by the new variables of the main func
		isValidName, isValidTicketNumber := checkInputValidity(firstName, lastName, userTickets)

		if isValidName && isValidTicketNumber {
			//processBooking function does the booking process and returns the value of new remaining tickets to further use in the main func
			newRemainingTickets := processBooking(firstName, userTickets, remainingTickets, nickNames)

			if newRemainingTickets == 0 {
				fmt.Println("SOLD OUT!!!!! No more tickets left for sale")

				break
			}
		} else {
			if !isValidName {
				fmt.Printf("check the length of the first name and last name\n")
			} else if !isValidTicketNumber {
				fmt.Printf("check the ticket count\n")

			}

		}
	}

}
func welcomeMessage(toolName string, totalTickets int, remainingTickets int) {
	fmt.Printf("Welcome to the %s ticket booking system\n", toolName) //use of functions defined outside the main func and then called within the func main
	fmt.Printf("We have a total of %d for sale\n", totalTickets)
	fmt.Printf("However only %d are available as of now\n", remainingTickets)
}

func gatherUserInput() (string, string, int) {
	fmt.Println("what is your first name?:") //this prints on the screen
	fmt.Scan(&firstName)
	fmt.Println("what is your last name")
	fmt.Scan(&lastName) //this will store the user input
	fmt.Printf("Hi %s, How many tickets do you need?\n", firstName)
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}

func checkInputValidity(firstName string, lastName string, userTickets int) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 //#we are preparing two new variables which will check for the conditions related to user input, this is further used by if statements in the main func
	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidTicketNumber
}

func processBooking(firstName string, userTickets int, remainingTickets int, nickNames []string) int {
	remainingTickets -= userTickets //new value of the variable

	//create map for user data
	var userInfo = make(map[string]string)
	userInfo["firstName"] = firstName
	userInfo["lastName"] = lastName
	userInfo["userTickets"] = strconv.FormatInt(int64(userTickets), 10)

	fmt.Printf("%s has purchased %d of tickets\n", firstName, userTickets)
	fmt.Printf("Thank you %s for your Purchase. A total of %d tickets will be sent to your email\n", firstName, userTickets)

	bookingsName = append(bookingsName, userInfo)

	for _, booking := range bookingsName { //_ is used for unused variables unless go keeps complaining
		nickNames = append(nickNames, booking["firstName"]) //here we are adding the firstName which is the nickName into the array nickNames (focus on the s used)
	}

	fmt.Printf("Remaining tickets are %d\n", remainingTickets)
	fmt.Printf("List of all the bookings %v \n", bookingsName)
	return remainingTickets

}
