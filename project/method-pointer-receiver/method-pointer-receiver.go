package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// secara default/bawaan properti pada method struct yang diakses didalam block kode method tersebut merupakan duplikasi nya
// jadi jika melakukan mutasi/perubahan terhadap nilai dari field/properti tidak berpengaruh atau terjadi perubahan pada struct
func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}

// oleh karena itu method harus menggunakan pointer receiver agar dapat melakukan perubahan terhadap nilai dari field sebuah struct
// method dari struct Person dengan pointer receiver berfungsi untuk melakukan perubahan pada nilai field Age dari nilai parameter
// yang di input
func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) SetName(name string) {
	p.Name = name
}

func main() {
	p1 := Person{Name: "Otong", Age: 19}

	fmt.Printf("%#v\n", p1)
	p1.SetAge(17)
	fmt.Printf("%#v\n", p1)

	p1.SetName("Ucup")
	// p1.Name tidak berubah menjadi "Ucup" dan nilai nya tetap "Otong" karena method tidak menggunakan pointer receiver
	// sehingga field Name pada method adalah duplikasi bukan nilai asli dari p1
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%s\n", p1.GetName())
}
