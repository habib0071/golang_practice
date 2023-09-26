package main

import "fmt"

func main() {
	fmt.Println("My name is habib")

	var WhatToSay string
	var i int
	var str string

	WhatToSay = "habib"
	fmt.Println(WhatToSay)

	i = 12

	fmt.Print("the number is:",i)
	fmt.Println(122)

	str = SaySomething()
	fmt.Println(str)

	str1 := SaySomething()

	fmt.Println(str1)

	s1, s2 := doubleString()

	fmt.Println(s1, s2)
}

func SaySomething() string {
	return "somthing"
}


func doubleString() (string, string) {
	return "Habib", "Nothing"
}
