package pkgaccessmodifier

import "fmt"

// ini adalah package untuk mencoba jika bisa membatasi sebuah akses untuk variabel, fungsi dan yang lainnya
// berdasarkan dari namanya yang diawali huruf besar atau kecil pada package yang berbeda

const (
	Hello = "Hello"
	hi    = "Hi" // variabel ini tidak akan bisa diakses oleh package lain
)

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayHi(name string) {
	fmt.Println("Hi", name)
}
