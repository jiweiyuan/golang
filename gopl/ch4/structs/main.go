package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

type Family struct {
	Man struct {
		Name string
	}

	Women struct {
		Name string
	}
}

func main() {
	var dilbert Employee
	dilbert.Salary += 50000
	fmt.Println(dilbert.Salary)

	position := &dilbert.Position
	*position += "Being"
	fmt.Println(dilbert.Position)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " XiCheng"
	fmt.Println(dilbert.Position)

	family := Family{
		Man: struct{ Name string }{Name: "Yuan"},
		Women: struct{ Name string }{Name: "Liu"},
	}
	fmt.Println(family.Man.Name)
}
