package main

import "fmt"

type Point struct {
	x float32
	y float32
}

type Person struct {
	firstName string
	lastName string
	age int
}

func (p *Person) updateNameFromPointer(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	delphi := Person{
		firstName: "Delphi",
		lastName: "Kuo",
		age: 31,
	}

	fmt.Println(delphi)

	delphi.updateName("Dolph")

	fmt.Println(delphi)

	delphi.updateNameFromPointer("Dow")

	fmt.Println(delphi)

}