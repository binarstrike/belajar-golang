package person

// jika ingin menambahkan lebih dari satu tag pada struct bisa dengan memisahkannya dengan spasi

type Person struct {
	FirstName string `json:"first_name" required:"true"`
	LastName  string `json:"last_name" required:"false"`
	Age       int    `json:"age" required:"true"`
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) SetLastName(lastName string) {
	p.LastName = lastName
}

func (p *Person) SetAge(age int) {
	p.Age = age
}
