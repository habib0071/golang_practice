package main

import (
	"log"
	"time"
)
// var firstname string
// var lastname string
// var phoneNumber string
// var age int
// var birthday time.Time

type User struct {
	firstname string
    lastname string
	phoneNumber string
	age int
	birthday time.Time
}

func main() {
	user := User {
		firstname: "Habib", 
		lastname: "Khan",
		phoneNumber: "1 555 123 12",


	}
	log.Println(user.firstname)
}