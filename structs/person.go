package structs

import "fmt"

type Person struct {
	// private fields (encapsulation)
	firstName string
	lastName  string
	age       int
}

func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

// Set Methods
func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}

func (p *Person) SetLastName(lastName string) {
	p.lastName = lastName
}

func (p *Person) SetAge(age int) {
	p.age = age
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
