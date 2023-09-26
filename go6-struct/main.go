package main

import "log"

type mystruct struct {
	FirstName string
}

func (m *mystruct) printname() string {
	return m.FirstName
}

func main() {
	var myVar mystruct

	myVar.FirstName = "John"

	myVar2 := mystruct{
		FirstName : "alita",
	}

	log.Println(myVar2.printname())
}