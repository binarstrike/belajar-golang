package main

import (
	"embed"
	"fmt"

	"github.com/goccy/go-json"
)

// package embed merupakan package yang digunakan untuk memasukan atau menambahkan
// sebuah atau banyak file kedalam file binary atau hasil keluaran dari proses build
// aplikasi.

// dengan menggunakan fitur ini dimungkinkan untuk menambahkan file kedalam binary
// golang hasil build pada project aplikasi seperti misal aplikasi web yang
// membutuhkan file statis dan template html

// untuk menggunakan fitur ini cukup dengan menulis comment pada baris sebelum deklarasi.
// untuk tipe data nya dapat berupa slice of byte atau string dan jika berupa folder maka
// dapat menggunakan tipe embed.FS
/*
//go:embed <nama file atau folder>
var file []byte

//go:embed static/*.png
var images embed.FS
*/
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//go:embed file.json
var people []byte

//go:embed folder1/*.txt
var folder1 embed.FS

//go:embed folder2
var folder2 embed.FS

func main() {
	var p []Person
	json.Unmarshal(people, &p)
	fmt.Printf("%#v\n", p)

	file1, _ := folder1.ReadFile("folder1/file1.txt")
	file2, _ := folder1.ReadFile("folder1/file2.txt")

	fmt.Println(string(file1), string(file2))

	dir, _ := folder2.ReadDir("folder2")

	for _, v := range dir {
		if v.IsDir() {
			fmt.Printf("dir folder2/%s\n", v.Name())
			continue
		}
		fmt.Printf("file folder2/%s\n", v.Name())
	}

	fmt.Println()

	txtfile, _ := folder2.ReadFile("folder2/txt/file1.txt")
	fmt.Println(string(txtfile))
}
