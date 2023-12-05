package main

import (
	"fmt"
)

// instal stringer terlebih dahulu
// go install golang.org/x/tools/cmd/stringer@latest

// stringer adalah sebuah perintah untuk menghasilkan file golang dengan isi
// berupa implementasi dari interface Stringer yaitu method String()
// dimana method ini akan digunakan untuk mengubah data dengan tipe selain string
// menjadi string normal yang dapat digunakan oleh kode program pada aplikasi

// ada beberapa parameter yang dapat digunakan pada perintah stringer yaitu
// -type untuk menentukan tipe data mana yang akan mengimplementasikan interface Stringer
// -output menentukan nama output file

//go:generate stringer -type=Foo -output=foo_string.go
type Foo int

const (
	foo1 Foo = iota
	foo2
	foo3
	bar1
	bar2
)

func main() {
	fmt.Println(foo1)          // foo1
	fmt.Println(bar2)          // bar2
	fmt.Println(foo3.String()) // foo3
}
