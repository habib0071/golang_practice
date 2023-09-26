package main 

import "log"

func main() {
	var myString  string

	myString = "Green"

	log.Println("myString is set to : ", myString)
	ChangeMyString(&myString)
	log.Println("Now, myString is set to : ", myString)
}

func ChangeMyString(s *string) {
	newValue := "red"

	*s = newValue
}