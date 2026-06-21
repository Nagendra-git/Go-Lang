package main

import (
	"fmt"
	"math"
)

// define circle struct
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	weidth, height float64
}

type value int

type student struct{
	grade string
}

//define a method for the circle

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rect Rectangle) area() float64 {
	return rect.height * rect.weidth
}

func (v value) cube() value {
	return v * v * v
}

func (s *student) updateGrade(value string){
	s.grade = value
}

func main() {

	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area : %f\n", circle.area())

	rectObj := Rectangle{weidth: 2.4, height: 4.5}
	fmt.Println("Area of Rectangle :", rectObj.area())

	x := value(3)
	y := x.cube()
	fmt.Println("Cube of", x, "is", y)

	s := student{grade: "A"}
	fmt.Println("Before:", s.grade)
	s.updateGrade("B")
	fmt.Println("After :", s.grade)


}
