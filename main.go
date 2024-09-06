package main

import (
	"fmt"
	"strings"
)

var ConferenceName= "Go conference"
const conferenceTickets int =50
var remainingTickets=50 
var bookings []string

func main(){
	 
	//bookings:= []string{}
	greetUsers()
	
	for remainingTickets>0 {
		firstName,lastName,email,userTickets:=getUserInputs()

		isValidName,isValidEmail,isValidTicketNumber:=validateUserInput(firstName,lastName,email,userTickets,remainingTickets)
		
		if  isValidName&&isValidEmail&&isValidTicketNumber{
			bookTickets(userTickets,firstName,lastName,email)
			

			firstNames:= getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n",firstNames)	
			noTickets:=remainingTickets==0
			if noTickets {
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
			}
		}else {
			if !isValidName{
				fmt.Println("first or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Email address you entered does not contain a @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid")
			}
			//fmt.Println("Your input data is invalid please try again")
			continue
		}
		
		
	}

	// city:="london"

	// switch city {
	// case "New York":

	// case "Singapore","Hong Kong":

	// case "London":
		
	// case "Berlin":

	// default:
	// 	fmt.Println("Invalid city")
	
	// }
	


}

func greetUsers(){
	
	fmt.Printf("Welcome to %v booking application\n",ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets) 
	fmt.Println("Get you tickets to attend")
}

func getFirstNames() []string{
	var firstNames []string
	for _,booking := range bookings{
		var names=strings.Fields(booking)
		firstNames=append(firstNames, names[0])
	}
	return firstNames
}

func getUserInputs() (string ,string, string, int){
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email Address: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName,lastName,email,userTickets
}

func bookTickets(userTickets int,firstName string,lastName string,email string){
	remainingTickets-=userTickets
	bookings=append(bookings,firstName+" "+lastName)

	fmt.Printf("The whole array: %v\n",bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("Remaining Tickets: %v for %v\n",remainingTickets,ConferenceName)
}