package person

type Person struct {
	FirstName, LastName string
	Age                 int
}

func (p *Person) SetPersonFirstName(firstName string) {
	p.FirstName = firstName
}
func (p *Person) SetPersonLastName(lastName string) {
	p.LastName = lastName
}
func (p *Person) SetPersonAge(age int) {
	p.Age = age
}
