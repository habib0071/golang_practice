package main

import (
	"log"
	"math"
)

const s string = "constant"
func main() {
	var a = "Habib"
	log.Println(a)

	log.Println(s)

	var b, c int = 1, 2
	log.Println(b, c)

	var d = true
	log.Println(d)

	var e int 
	log.Println(e)

	f := "apple"
	log.Println(f)

	const n = 5000000000000
	const p = 3e20/n
	log.Println(int64(p))

	log.Println(math.Sin(2))


}