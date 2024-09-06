package main

import (
	"fmt"
	"strings"
)

func main(){
	var ConferenceName= "Go conference"
	const conferenceTickets=50
	var remainingTickets=50 
	fmt.Printf("Welcome to %v booking application\n",ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets) 
	fmt.Println("Get you tickets to attend")

	var bookings []string 

	
	for remainingTickets>0 {
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
		if userTickets<=remainingTickets {
			remainingTickets-=userTickets
			bookings=append(bookings,firstName+" "+lastName)

			fmt.Printf("The whole array: %v\n",bookings)

			fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
			fmt.Printf("Remaining Tickets: %v\n",remainingTickets)

			var firstNames []string

			for _,booking := range bookings{
				var names=strings.Fields(booking)
				firstNames=append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n",firstNames)

			noTickets:=remainingTickets==0
			if noTickets {
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
			}
		}else {
			fmt.Printf("We only have %v tickets remaining, so you cannot book %v tickets",remainingTickets,userTickets)
			continue
		}
		
		
	}

	
	


}