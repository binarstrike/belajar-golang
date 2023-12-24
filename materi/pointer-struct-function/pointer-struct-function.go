package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

// secara default instance struct Person yang memakai method atau fungsi SetPersonAge_ adalah
// duplikat dari struct tersebut jadi Person.Age tidak akan ikut berubah,
func (p Person) SetPersonAge_(age int) {
	p.Age = age
}

// agar untuk itu gunakan pointer pada deklarasi fungsi pada struct.
func (p *Person) SetPersonAge(age int) {
	p.Age = age
}

func (p *Person) SetPersonFirstName(firstName string) {
	p.FirstName = firstName
}

func main() {
	p1 := Person{FirstName: "Ucup", LastName: "Surucup", Age: 18}
	p2 := &p1

	fmt.Printf("p2: %#v\n", p2)

	p2.SetPersonAge_(20) // p2.Age tidak berubah nilainya
	fmt.Printf("p2: %#v\n\n", p2)

	p2.SetPersonAge(19)
	fmt.Printf("p2: %#v\n", p2)
	p2.SetPersonFirstName("Budi")
	fmt.Printf("p2: %#v\n", p2)
}
