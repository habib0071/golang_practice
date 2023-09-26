/*package main

import (
	"log"
	"math"
)
type caluclation interface {
	area() float64
	perimeter() float64
}

type rect struct {
	Width, Hight float64

}
type circle struct {
	R float64
}


func main() {
	rect1 := rect{
		Width: 12,
		Hight: 10,
	}
	PrintSomething(rect1)

	circle1 := circle{
		R: 4,
	}
	PrintSomething(circle1)
}

func PrintSomething(a caluclation) {
	log.Println("the area will be: ", a.area(), "and perimeter will be: ", a.perimeter())
}

func (r rect) area() float64 {
	return r.Hight * r.Width
}
func(r rect) perimeter() float64 {
	return 2*r.Width + 2*r.Hight
}

func (c circle) area() float64 {
	return math.Pi * c.R * c.R
}
func(c circle) perimeter() float64 {
	return 2 * math.Pi * c.R
}*/

package main

import (
	"github.com/habib0071/golang_practice/go13-package/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	myVar.TypeNumber = 12
	log.Println(myVar)

}