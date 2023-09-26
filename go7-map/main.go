package main

import "log"

func main() {
	var myMap0 map[string]string
	myMap := make(map[string]string)
	myMap0 = make(map[string]string)

	myMap["dog"] = "kitty"
	myMap["cat"] = "tommy"
	myMap0["name"] = "Habib"

	log.Println(myMap["dog"])
	log.Println(myMap0["name"])

	
}