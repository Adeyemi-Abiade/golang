package main

import "fmt";

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets 

	// Printing out the types of variables
	fmt.Printf("conferenceName is of type %T\n", conferenceName)
	fmt.Printf("conferenceTickets is of type %T\n", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	//ask user for their input
	fmt.Print("Username: ")
	fmt.Scan(&userName)
	fmt.Print("No of tickets: ")
	fmt.Scan(&userTickets)
	fmt.Printf("User, %v, ordered %v number of tickets\n", userName, userTickets)
}
