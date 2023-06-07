package main

import "fmt"

//here we try third principle of SOLID:
//Liskov substitution principle - LSP
//LSP is about inheritance or composition
//You can use the same methods, if struct realize another struct with this method
//What is wanted here is something like the following substitution property:
//If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T,
//the behavior of P is unchanged when o1 is substituted for o2 then S is a subtype of T

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
