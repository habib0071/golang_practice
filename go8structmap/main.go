/*package main

import "log"

type User struct{
	FirstName string
	LastName string
}

func main() {

	user1 := make(map[string]User)

	me := User {
		FirstName: "johon",
		LastName: "abrahaom",

	}
	user1["me"] = me

	log.Println(user1["me"].FirstName)
}*/

/*package main

import "log"

type User struct {
	FirstName string
	LastName string
	phoneNumber string
	age int
	FatherName string
	MontherName string
}

func main() {
	User2 := make(map[string]User)

	me := User {
		FirstName: "Habib",
		LastName: "Khan",
		phoneNumber: "01812929409",
		age: 24,
		FatherName: "Ruhul amin",
		MontherName:"Halima begum",
	}

	User2["u"] = me

	me1 := User {
		FirstName: "Ohona",
		LastName: "Khan",
		phoneNumber: "01812929409",
		age: 16,
		FatherName: "Ruhul amin",
		MontherName:"Halima begum",
	}
	User2["u1"] = me1

	log.Println("Full Name: ", User2["u"].FirstName, User2["u"].LastName)
	log.Println("Phone Number: ", User2["u"].phoneNumber)
	log.Println("Age: ", User2["u"].age)
	log.Println("Father Name: ", User2["u"].FatherName)
	log.Println("Monther Name: ", User2["u"].MontherName)

	log.Println()

	log.Println("Full Name: ", User2["u1"].FirstName, User2["u"].LastName)
	log.Println("Phone Number: ", User2["u"].phoneNumber)
	log.Println("Age: ", User2["u"].age)
	log.Println("Father Name: ", User2["u"].FatherName)
	log.Println("Monther Name: ", User2["u"].MontherName)

	log.Println()

	log.Println(User2["u"], User2["u1"])
}*/

package main

import (
	"log"
	"sort"
)

func main() {
	var myVar float32

	myVar = 12.34
	log.Println(myVar)

	var myString []string
	
	myString = append(myString, "habib")
	myString = append(myString, "ohona")

	log.Println(myString[1])

	var myIntList []int

	myIntList = append(myIntList, 3)
	myIntList = append(myIntList, 1)
	myIntList = append(myIntList, 7)
	myIntList = append(myIntList, 5)

	log.Println(myIntList)

	sort.Ints(myIntList)

	log.Println(myIntList)

	number := []int {1, 2, 3, 4, 5, 6}

	log.Println(number)
	log.Println(number[1:3])

	name := []string {"one", "two", "three", "four"}
	log.Println(name)
}
