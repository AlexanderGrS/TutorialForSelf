package main

import "fmt"

//And this is last - fifth principe
//Dependency Inversion Principle - DIP
//Objects should depends on abstractions, and not on another objects
//High-level modules should not depend on low-level modules.
//Both should depend on abstractions.
//Abstractions should not depend on details.
//Details should not depend on abstractions

type DBConn interface {
	QuerySomeData() []string
}

type MySQL struct {
}

func (db MySQL) QuerySomeData() []string {
	return []string{"obj1", "obj2", "obj3"}
}

type MyRepository struct {
	db DBConn // dependency on abstraction
	//db MySQL - dependency on object
}

func (r MyRepository) GetData() []string {
	return r.db.QuerySomeData()
}

func main() {
	mysqlDB := MySQL{}
	repo := MyRepository{db: mysqlDB}
	fmt.Println(repo.GetData())
}
