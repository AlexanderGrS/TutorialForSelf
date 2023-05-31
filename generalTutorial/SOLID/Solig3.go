package main

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

func (cube) area() float64 {

}
