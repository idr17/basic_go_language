package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstname string
	// lastname  string
	// gender    string
	// age       int

	firstname, lastname, gender string
	age                         int
}

// greeting method (value receiver). biasanya value receiver digunakan untuk calculation/perhitungan
func (p Person) greet() string {
	return "Hello, my name is " + p.firstname + " " + p.lastname + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver). untu meruba value gunakan pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

// hasMarried method (pointer receiver)
func (p *Person) hasMarried(spousename string) {
	if p.gender == "f" {
		p.lastname = spousename
	}
}

func main() {
	person1 := Person{firstname: "Indro", lastname: "Allezz", age: 19, gender: "m"}
	person2 := Person{firstname: "Ana", lastname: "Maulida", age: 18, gender: "f"}

	person1.hasBirthday()
	person1.hasMarried("ana")

	person2.hasMarried("Allezzii")
	person2.hasBirthday()

	fmt.Println(person2.greet())
}
