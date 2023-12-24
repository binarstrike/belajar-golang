package main

// Menurut ChatGPT:
// Access modifier dalam pemrograman adalah istilah yang digunakan untuk mengatur tingkat akses ke anggota
// (variables, functions, classes, dll) dalam sebuah program. Ini memungkinkan pengembang untuk mengontrol
// siapa yang bisa melihat atau mengubah suatu anggota. Terdapat beberapa jenis access modifier seperti "public"
// (dapat diakses dari mana saja), "private" (hanya dapat diakses dari dalam kelas itu sendiri), dan "protected"
// (hanya dapat diakses dalam kelas yang sama atau kelas turunannya).

// Dalam bahasa pemrograman Go (Golang), konsep aksesibilitas atau "access modifier" berbeda dari beberapa bahasa
// pemrograman lain. Go memiliki aturan yang sederhana: jika sebuah identifier (variabel, fungsi, atau struktur data)
// diawali dengan huruf besar (seperti NamaVariabel), itu bersifat publik dan dapat diakses dari
// luar paket (package). Jika diawali dengan huruf kecil (seperti namaVariabel), itu bersifat privat dan hanya
// dapat diakses di dalam paket yang sama.

// jadi hak akses dari sebuah variabel atau fungsi dan lainnya pada package yang berbeda ditentukan dari nama nya
// jika diawali huruf besar maka itu public bisa diakases siapa saja dan jika diawali huruf kecil maka itu
// private dan hanya bisa diakses oleh package yang membuat atau mendeklarasikan nya

import (
	"fmt"

	pkgaccessmodifier "github.com/binarstrike/belajar-golang/pkg/pkg-access-modifier"
)

func main() {
	fmt.Println(pkgaccessmodifier.Hello, "Udin")
	// fmt.Println(coba.hi, "Ucup") // ini akan error karena varibel diawali huruf kecil

	pkgaccessmodifier.SayHello("Budi")
	// coba.sayHi("Otong") // error
}
