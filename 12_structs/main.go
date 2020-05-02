package main

import (
	"fmt"
	"strconv"
)

type Person struct { // Define person struct
	/* firstName string
	lastName  string
	city      string
	gender    string
	age       int */

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver) { p is like this on other PL.}
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBday method (pointer receiver)
func (p *Person) hasBday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastname
	}
}

func main() {
	// init person struct
	person1 := Person{firstName: "Sasa", lastName: "Meow", city: "Earth", gender: "F", age: 11}
	person2 := Person{"Bob", "Johnson", "Earth", "M", 33}
	fmt.Println(person1)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	// Call hasBday method
	person1.hasBday()
	person1.getMarried("MeMeow")

	person2.getMarried("Thompson")
	// CALL Greeting method
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
