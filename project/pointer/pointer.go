package main

import (
	"fmt"
)

// Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
// Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
// sebenarnya yang dikirim adalah duplikasi value nya.

// Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama,
// tanpa menduplikasi data yang sudah ada
// Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

type Person struct {
	FirstName, LastName string
	Age                 int
}

func pointer() {
	p1 := Person{
		FirstName: "Binar",
		LastName:  "Nugroho",
		Age:       18,
	}

	p2 := p1    // p2 adalah duplikasi dari p1
	p2.Age = 19 // p1.Age tidak akan berubah nilainya
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p2: %#v\n", p2)

	fmt.Println()

	// Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, bisa menggunakan
	// operator & diikuti dengan nama variable nya
	p3 := &p1               // p3 pointer ke p1
	p3.FirstName = "Albert" // p1.FirstName nilainya akan berubah menjadi "Albert"
	p3.LastName = "Ucup"    // p1.LastName nilainya akan berubah menjadi "Ucup"
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p3: %#v\n", p3)
}

func menetapkanPointerBaruPadaPointer() {
	person1 := Person{FirstName: "Dika", LastName: "Purnama", Age: 18}
	person2 := &person1

	fmt.Printf("\nperson1: %#v\n", person1)
	person2.FirstName = "Andi" // person1.FirstName -> "Andi"
	fmt.Printf("person1: %#v\n", person1)
	fmt.Printf("person2: %#v\n\n", person2)

	person2 = &Person{FirstName: "Gilang", LastName: "Foo", Age: 19}
	// person2 tidak lagi menjadi pointer ke person1 dan membuat pointer ke struct Person dengan alamat baru
	person1.Age = 20
	fmt.Printf("person1: %#v\n", person1)
	fmt.Printf("person2: %#v\n", person2)
}

func mengubahSeluruhDataPadaSatuPointer() {
	orang1 := Person{FirstName: "Binar", LastName: "Nugroho", Age: 18}
	orang2 := &orang1
	orang3 := &orang1

	fmt.Printf("\norang1: %#v\n", orang1)
	orang2.Age = 17 // orang1.Age -> 17
	fmt.Printf("orang1: %#v\n", orang1)
	fmt.Printf("orang2: %#v\n", orang2)

	fmt.Println()

	// untuk mengubah data agar semua yang mengacu ke pointer yang sama ikut berubah gunakan karakter `*`
	// di depan nama variabel dan assign dengan data baru
	*orang2 = Person{FirstName: "Bimo", LastName: "Anandika", Age: 18}
	// mengakses pointer orang2 dan mengubah semua variabel atau instance yang satu reference atau alamat dengan orang2
	// dengan struct Person baru
	fmt.Printf("orang1: %#v\n", orang1)
	fmt.Printf("orang2: %#v\n", orang2)
	fmt.Printf("orang3: %#v\n", orang3)
}

func pointerPadaFunction() {
	// secara default nilai apapun yang di pasing ke parameter fungsi adalah duplikat dari nilai sebenarnya
	// jadi variabel yang di passing ke parameter tersebut tidak akan berubah nilainya.
	// untuk melakukan perubahan pada variabel atau nilai yang di passing ke parameter fungsi caranya dengan
	// menambahkan * pada tipe data parameter

	// fungsi setPersonFirstName adalah fungsi untuk mengubah nilai FirstName pada struct Person
	// Parameter p adalah pointer ke Person, sehingga perubahan yang dilakukan di sini akan mempengaruhi nilai asli.
	setPersonFirstName := func(p *Person, firstname string) {
		p.FirstName = firstname
	}
	// fungsi setPersonAge adalah fungsi untuk mengubah nilai Age pada struct Person dan sama seperti fungsi
	// setPersonFirstName parameter p adalah pointer ke struct Person
	setPersonAge := func(p *Person, age int) {
		p.Age = age
	}

	p1 := Person{FirstName: "Bagas", Age: 18}
	p2 := &p1

	fmt.Printf("p1: %#v\n", p1)
	// jika nilai atau variabel yang di passing bukan sebuah pointer maka harus menambahkan karakter &
	setPersonFirstName(&p1, "Budi")
	setPersonAge(p2, 19)
	fmt.Printf("p1: %#v\n", p1)
}

func main() {
	pointer()
	fmt.Println("============================================")
	menetapkanPointerBaruPadaPointer()
	fmt.Println("============================================")
	mengubahSeluruhDataPadaSatuPointer()
	fmt.Println("============================================")
	pointerPadaFunction()
}
