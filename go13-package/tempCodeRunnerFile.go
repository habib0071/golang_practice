package main

import "log"

type caluclation interface {
	area() int
	perimeter() int
}

type rect struct {
	width, hight int

}
type circle struct {
	r int
}


func main() {
	rect1 := rect{
		width: 12,
		hight: 10,
	}
	PrintSomthing(rect1)
}

func PrintSomthing(a caluclation) {
	log.Println("the area will be: ", a.area(), "and perimeter will be: ", a.perimeter())
}

func (r rect) area() int {
	return r.width * r.hight
}