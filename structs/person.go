package structs

import (
	"errors"
	"fmt"
)

const MIN_NAME_LENGTH = 2

type Person struct {
	// private fields (encapsulation)
	firstName string
	lastName  string
	age       byte //alias for type uint8 (just to know)
}

// Constructor
func NewPerson(firstName, lastName string, age byte) (*Person, error) {
	if len(firstName) < MIN_NAME_LENGTH || len(lastName) < MIN_NAME_LENGTH {
		return nil, errors.New("\nerror: invalid name\nmsg: first and last name must have at least 2 characters")
	}

	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}, nil
}

// Setters
func (p *Person) SetFirstName(firstName string) error {
	if len(firstName) < 2 {
		return errors.New("\nerror: invalid name\nmsg: first name must have at least 2 characters")
	}

	p.firstName = firstName
	return nil
}

func (p *Person) SetLastName(lastName string) error {
	if len(lastName) < 2 {
		return errors.New("\nerror: invalid name\nmsg: last name must have at least 2 characters")
	}

	p.lastName = lastName
	return nil
}

func (p *Person) SetAge(age byte) {
	p.age = age
}

// Getters
func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) Age() byte {
	return p.age
}

// Other
func (p *Person) HaveABirthday() {
	p.age += 1
	fmt.Printf("%s is now %d years old.\n", p.firstName, p.age)
}
