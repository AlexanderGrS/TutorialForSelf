package main

import (
	"fmt"
	"math"
)

//This project is about two first principles of SOLID:
//Single Resposibility Principle - S
//Open-Close Principle - O
//SRP is about separation of calculating area() and output areaToText()
//A module should have one, and only one reason to change
//OCP is about using of any type for one function(not choosing between limited types), that implements the same methods
//A software artifact should be open for extencion but closed for modifications

type calculateSquareOfArea interface {
	area() float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	sideLen float32
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

type trinagle struct {
	height float32
	base   float32
}

func (t trinagle) area() float32 {
	return t.base * t.height / 2
}

type calculator struct{}

func (c calculator) sumOfAreas(areas ...calculateSquareOfArea) (sum float32) {
	for _, area := range areas {
		sum += area.area()
	}
	return
}

type outPrinter struct{}

func (outPrinter) areaToText(c calculateSquareOfArea) string {
	return fmt.Sprintf("the area is: %f", c.area())
}

func main() {
	c := circle{radius: 7}
	s := square{sideLen: 12}
	t := trinagle{height: 10, base: 15}

	out := outPrinter{}
	calculator := calculator{}
	fmt.Println(out.areaToText(c))
	fmt.Println(out.areaToText(s))
	fmt.Println(out.areaToText(t))
	fmt.Println(calculator.sumOfAreas(c, s, t))

}
