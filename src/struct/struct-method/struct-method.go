package main

import "fmt"

/*
Menurut ChatGPT:
Struct method dalam Go adalah fungsi yang terkait dengan suatu tipe data struct.
Fungsi ini beroperasi pada data yang dimiliki oleh struct tersebut. Ini memungkinkan kamu untuk
menambahkan perilaku atau tindakan khusus yang dapat dilakukan oleh instance dari struct tersebut.
*/

// struct pada golang juga dapat menyimpan method atau fungsi. Instance atau variabel dari tipe data
// struct ini dapat mengakses dan mejalankan fungsi tersebut seperti mengakses field pada struct
type Person struct {
	FirstName, LastName string
}

type Persegi struct {
	Sisi int
}

func (p Person) SayHello() {
	fmt.Println("Hello", p.FirstName)
}

func (persegi Persegi) HitungKeliling() int {
	return persegi.Sisi * 4
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

func main() {
	p1 := Person{
		FirstName: "Binar",
		LastName:  "Nugroho",
	}
	p1.SayHello()

	fmt.Println()

	persegi1 := Persegi{
		Sisi: 10,
	}

	fmt.Println("Keliling persegi1:", persegi1.HitungKeliling())
	fmt.Println("Luas persegi1:", persegi1.HitungLuas())
}
