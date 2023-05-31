package main

import "fmt"

type Person struct {
	Name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	pers1.Name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	pers2.Name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	printPerson(pers1)
	printPerson(pers2)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.Name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
