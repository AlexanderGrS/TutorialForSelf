package main

import (
	"fmt"
	"math"
)

//This is fourth SOLID principe
//Interface Segregation Principle  - ISP
//It is about architecture, where we don't use
//methods, that are not used by structs
//Clients should not be forced to depend on methods they do not use

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

type square struct {
	sideLen float64
}

func (s square) area() float64 {
	return s.sideLen * s.sideLen
}

type cube struct {
	square
}

func (c cube) area() float64 {
	return math.Pow(c.sideLen, 2)
}

func (c cube) volume() float64 {
	return math.Pow(c.sideLen, 3)
}

func areaSum(shapes ...shape) (sum float64) {
	for _, s := range shapes {
		sum += s.area()
	}
	return
}

func areaVolumeSum(objects ...object) (sum float64) {
	for _, o := range objects {
		sum += o.area() + o.volume()
	}
	return
}

func main() {
	s := square{4}
	c := cube{s}

	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(c))
}
