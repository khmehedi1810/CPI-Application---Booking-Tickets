package main

import (
	"fmt"
	"go/help"
	"sync"
	"time"
)

var info1 = "WelCome to our conference booking application"
var info2 = "TO attent the event you have to confirm a ticket"
var uptik uint = 50
var booking = make([]userInfo, 0) //size declare kore dite hobe

type userInfo struct {
	fname string
	lname string
	email string
	notik uint
}

var wp = sync.WaitGroup{}

func main() {

	info()

	fname, lname, email, tik := getUserin()

	//validate user input function
	namVal, emVal, tickVal := help.UserinVal(fname, lname, email, tik, uptik)

	if namVal && emVal && tickVal {
		bookin(tik, fname, lname, email)
		wp.Add(1) //as there are 1 go. so I declare  1here, if there were two go, then i putted 2
		go sendMail(tik, fname, lname, email)

		//Call function for first name print
		firstNames := firNam()

		fmt.Printf("All Our Bookings: %v \n", firstNames)

		if uptik <= 0 {
			fmt.Println("All ticket is sold out. Better luck for the next time")

		}
	} else {
		if !namVal {
			fmt.Println("Your name is incorrect, input at least 2 characters")
		}
		if !emVal {
			fmt.Println("Your email is incorrect, input currect email")
		}
		if !tickVal {
			fmt.Println("Please Enter valid information again")
		}

	}
	wp.Wait()

}

func info() {
	fmt.Println(info1)
	fmt.Println(info2)
}

func firNam() []string {
	firstNames := []string{}
	for _, boo := range booking {
		firstNames = append(firstNames, boo.fname)
	}
	return firstNames
}

func getUserin() (string, string, string, uint) {
	var fname, lname, email string
	var tik uint
	fmt.Println("Enter First Name & Last Name: ")
	fmt.Scan(&fname, &lname)

	fmt.Println("Enter Email Address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of ticket that you want to book")
	fmt.Scan(&tik)

	return fname, lname, email, tik
}

func bookin(tik uint, fname string, lname string, email string) {
	uptik = uptik - tik

	var useDara = userInfo{
		fname: fname,
		lname: lname,
		email: email,
		notik: tik,
	}

	booking = append(booking, useDara)
	fmt.Printf("List is: %v\n", booking)

	fmt.Printf("Thanks %v %v for Confirming ticket. You have bought %d tickets\n", fname, lname, tik)
	fmt.Println("Remaining Tickets is: ", uptik)

}

func sendMail(tik uint, fname string, lname string, email string) {
	time.Sleep(10 * time.Second)
	var tikinfo = fmt.Sprintf("%v tickets for %v %v ", tik, fname, lname)
	fmt.Println("############")
	fmt.Printf("Sending tickets: %v\n to email address %v\n", tikinfo, email)
	fmt.Println("############")
	wp.Done()
}
