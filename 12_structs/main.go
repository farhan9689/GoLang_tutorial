package main

import (
	"fmt"
	"strconv"
)

// creating struct
type person struct {
	// name   string
	// age    int
	// gender string
	// city   string

	firstName, lastName, gender, city string
	age                               int
}

// greeting method (value reciever) -> example of value reciever method

func (p person) greet() string { // (p person)
	return "hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) //
}

// hasBirthday method (Pointer Reciever)
func (p *person) hasBirthday() {
	p.age++
}

//getMarried method (also a pointer reciever)

func (p *person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}

}

func main() {

	//initialising struct properties with values

	person1 := person{firstName: "farhan", lastName: "mirza", age: 24, gender: "m", city: "bangalore"}

	fmt.Println(person1)

	//Alternate way to initialise struct properties

	person2 := person{"mirza", "ashraf", "m", "nagpur", 25}
	person3 := person{"Aaliya", "khan", "f", "bangalore", 23}

	fmt.Println(person3)

	// getting single field

	//fmt.Println(person1.name)

	person1.hasBirthday()
	person2.hasBirthday()
	person3.getMarried("Mirza")
	fmt.Println(person3.greet())

}
