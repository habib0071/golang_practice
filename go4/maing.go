package main

import "log"

var s = "One"

func main() {
	var s1 = "two"
	log.Println(s)
	log.Println(s1)

	saySomthing("xxx")
}
func saySomthing(s3 string) (string, string) {
	log.Println(s, s3)
	return s, "world"

}