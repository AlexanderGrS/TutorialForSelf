package main

import "fmt"

type transport interface {
	printName() string
}

type vehicle struct {
	name string
}

func (v vehicle) printName() string {
	return v.name
}

type car struct {
	vehicle
	wheel int
	gates int
}

type motocycle struct {
	vehicle
	wheels int
}

type printer struct{}

func (printer) printTransportName(p transport) string {
	return fmt.Sprintf("Name: %s", p.printName())
}

func main() {
	v := vehicle{name: "Ford"}

	c := car{
		vehicle: vehicle{name: "Toyota"},
		wheel:   4,
		gates:   2,
	}

	m := motocycle{
		vehicle: vehicle{name: "Mitsibishi"},
		wheels:  2,
	}

	p := printer{}
	fmt.Println(p.printTransportName(c))
	fmt.Println(p.printTransportName(v))
	fmt.Println(p.printTransportName(m))
}
