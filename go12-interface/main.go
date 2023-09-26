package main

import "log"


type Animal interface {
	Say() string
	NumberOfLeg() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name: "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(dog)

	gorilla := Gorilla{
		Name: "Jack",
		Color: "black",
		NumberOfTeeth: 39,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("this animal says", a.Say(), "and has", a.NumberOfLeg(), "legs")
}

func (d Dog) Say() string {
	return "woof"
}

func (d Dog) NumberOfLeg() int {
	return 4
}

func (d *Gorilla) Say() string {
	return "ughhh"
}

func (d *Gorilla) NumberOfLeg() int {
	return 2
}