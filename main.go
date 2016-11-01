/*
Borrowing and stricts are what Frenando would like me to focus on here. Person is a basic person struct. Developer borrows from Person and adds the language perameter
*/

package main

import "fmt"

type Age int

func (a *Age) Birthday() {
	*a++
}

type Person struct {
	age   Age
	first string
	last  string
}

func (p Person) String() string {
	return fmt.Sprintf("Full name: %v, %v  Age: %v", p.first, p.last, p.age)
}

func (p *Person) Birthday() {
	person := *p
	person.age.Birthday()
	*p = person
}

type Developer struct {
	Person
	language string
}

func (d Developer) String() string {
	return fmt.Sprintf("%v, uses language: %v", d.Person.String(), d.language)
}

// Developers age twice as fast
func (d *Developer) Birthday() {
	dev := *d
	dev.Person.age.Birthday()
	dev.Person.age.Birthday()
	*d = dev
}

type Ager interface {
	Birthday()
}

func HardTimes(a Ager) {
	fmt.Println("Hard times have fallen...")
	a.Birthday()
	a.Birthday()
}

func main() {

	kory := Person{
		32,
		"Kory",
		"Donati",
	}
	fmt.Println(kory)
	kory.age.Birthday()
	fmt.Println(kory)
	HardTimes(&kory)
	fmt.Println(kory)

	fernando := &Developer{
		Person{
			33,
			"Fernando",
			"Sanchez",
		},
		"golang",
	}

	fmt.Println(fernando)
	fernando.Birthday()
	fmt.Println(fernando)
	HardTimes(fernando)
	fmt.Println(fernando)
}
