package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type PersonAgain struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := Person{
		firstName: "Jane",
		lastName:  "Doe",
		age:       30,
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age)

	p2 := Person{"Jane", "Doe", 30}

	fmt.Println(p2 == p1)

	// Can convert between types that have the same fields
	pAgain := PersonAgain(p1)
	fmt.Println(pAgain.firstName, "is", pAgain.age, "years old")
}

// Make two structs, Person and PersonAgain with the same fields. Convert between them
