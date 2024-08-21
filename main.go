package main

import "fmt"

func main()  {

	var confrenceName string = "Go confrence"
	const confrenceTicket int = 60
	var remainingTickets uint = 50
	var bookings = []string

	fmt.Printf("confrenceTicket is %T remainingTickets is %T, confrenceName is %T\n",confrenceName,confrenceTicket,remainingTickets )
	fmt.Printf("Welcome to %v booking application\n" , confrenceName)
	fmt.Println("we have total of", confrenceTicket, "tickets and", remainingTickets, "left" )
	fmt.Println("Get your tickets here to attend")



	for{
	 var firstname string
	var lastname string
	var email string
	var userTicket uint
	
	//ask user for their data
	fmt.Println("Enter your firstname ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your lastname ")
	fmt.Scan(&lastname)
	fmt.Println("Enter your email ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Println(&userTicket)

	remainingTickets= remainingTickets-userTicket
	bookings=append(bookings, firstName + " " +lastname) 

	fmt.Println("the whole array: %v\n",bookings)
	fmt.Println("the first value: %v\n",bookings[0])
	fmt.Println("the slice type: %T\n",bookings)
	fmt.Println("the slice length: %v\n",len(bookings) )

	
	fmt.Printf("Thank you %v %v for booking %v tickets. you will recieve a conformation mail at %v\n" , firstname,lastname,email,userTicket)
	fmt.Println("%v tickets remaing for %v\n", remainingTickets,confrenceName)

	}
	

	fmt.Println(remainingTickets)
	//print the memory address
	fmt.Println(&remainingTickets)

}

	
