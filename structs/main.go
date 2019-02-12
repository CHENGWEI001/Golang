package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gamil.com",
			zipCode: 94000,
		},
	}
	jim.updateName("jimmy")
	jim.print()
}

func (ponterToPerson *person) updateName(newFirstName string) {
	(*ponterToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
