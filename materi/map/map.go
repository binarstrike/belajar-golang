package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Sama seperti bahasa pemrograman lain golang juga memiliki fitur bernama Map
// fungsi nya sama seperti yang ada padsa bahasa javascript yaitu untuk menyimpan
// data dalam bentuk key-value(kata kunci dan nilai)
func main() {
	printMe := color.New(color.FgHiCyan)
	data := map[string]string{
		"nama":    "Binar Nugroho",
		"umur":    "12",
		"address": "Cilacap",
	}

	fmt.Println(data)

	// untuk mengakses data pada map dapat dilakukan dengan cara yang sama seperti array
	// namun untuk indeks nya menggunakan key atau kata kunci  map["key"]
	fmt.Println(data["nama"])
	fmt.Println(data["umur"])
	fmt.Println(data["address"])

	// untuk mengetahui panjang array bisa mneggunakan fungsi len(map)
	fmt.Printf("\nPanjang dari map %s adalah %d\n", "data", len(data))

	pemisah()

	// untuk membuat map baru tanpa harus langsung mengisi data nya bisa menggunakan fungsi make()
	post1 := make(map[string]string)
	post1["title"] = "Cara memasak air dengan benar"
	post1["description"] = "Tuang 40 liter ke dalam panci dan rebus dengan kompor"
	post1["wrong_data"] = "Blah"
	printMe.Printf("post1: %v\n", post1)
	printMe.Printf("panjang dari map post1: %d\n", len(post1))

	// untuk menghapus data yang ada pada map bisa menggunakan fungsi delete(map, key)
	delete(post1, "wrong_data")
	printMe.Printf("\npost1: %v\n", post1)
	printMe.Printf("panjang dari map post1: %d\n", len(post1))

	pemisah()

	names := make(map[int]string)
	names[1] = "Budi"

	// variabel kedua pada saat mengakses map adalah nilai boolean yang akan digunakan untuk pengecekan jika nilai yang
	// diakses pada map tidak kosong
	if name, ok := names[1]; ok {
		printMe.Printf("nama ditemukan pada key 1: %s\n", name)
	}

	if name, ok := names[5]; ok {
		printMe.Printf("nama ditemukan pada key 5: %s\n", name)
	} else {
		printMe.Println("nama tidak ditemukan pada key 5")
	}
}

func pemisah() {
	color.New(color.FgHiYellow).Println("=============================================================")
}
