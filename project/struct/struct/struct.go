package main

import (
	"fmt"
)

/*
Menurut ChatGPT:
Struct adalah tipe data di Go yang digunakan untuk menggabungkan beberapa nilai menjadi satu kesatuan.
Ini mirip dengan konsep "record" atau "struktur data" di bahasa pemrograman lainnya.
Struct digunakan untuk membuat tipe data kustom yang memiliki beberapa bidang atau properti.
Setiap bidang dalam struct dapat memiliki tipe data yang berbeda.
*/
type Person struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	// ada beberapa cara mendeklarasikan variabel dengan tipe data struct
	p1 := Person{
		FirstName: "Binar",
		LastName:  "Nugroho",
		Age:       18,
	}
	var p2 Person
	p2.FirstName = "Albert"
	p2.LastName = "Ucup"
	p2.Age = 21

	fmt.Printf("Person 1: %#v\n", p1)
	fmt.Printf("Person 2: %#v\n", p2)
}
