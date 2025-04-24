package structs

import (
	"errors"
	"fmt"
)

const (
	MIN_NAME_CHARS = 2
	MIN_AGE        = 0
)

type Person struct {
	// private fields (encapsulation)
	firstName string
	lastName  string
	age       int
}

func NewPerson(firstName, lastName string, age int) (*Person, error) {
	if len(firstName) < MIN_NAME_CHARS || len(lastName) < MIN_NAME_CHARS {
		return nil, errors.New("\nerror: invalid name\nmsg: first and last name must have at least 2 characters")
	}

	if age < MIN_AGE {
		return nil, errors.New("\nerror: invalid age\nmsg: must not be negative")
	}

	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}, nil
}

// Set Methods
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

func (p *Person) SetAge(age int) error {
	if age < 0 {
		return errors.New("\nerror: invalid age\nmsg: age must not be negative")
	}
	p.age = age
	return nil
}

// Get Methods
func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) GetLastName() string {
	return p.lastName
}

func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) GetAge() int {
	return p.age
}

// Other
func (p *Person) HaveABirthday() {
	p.age += 1
	fmt.Printf("%s is now %d years old.\n", p.firstName, p.age)
}
