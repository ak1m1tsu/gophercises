package structs

import "fmt"

type ContactInfo struct {
	Email   string
	ZipCode int
}

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

func NewPerson(firstName, lastName, email string, zipCode int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		ContactInfo: ContactInfo{
			Email:   email,
			ZipCode: zipCode,
		},
	}
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}

func (p *Person) UpdateName(newFirstName string) {
	p.FirstName = newFirstName
}
