package main

import "fmt"

// Generic juga bisa digunakan pada map

// type alias Foo dengan type parameter generic menjadikan map menjadi constraint atau tetapan bahwa key dari value pada map
// ini hanya dapat diberi nilai dengan tipe data int dan untuk nilainya dengan tipe data string
type Foo[K int, V string] map[K]V

func main() {
	foo := Foo[int, string]{
		1: "Bar",
		2: "Bar",
		3: "Bar",
	}

	fmt.Printf("foo: %#v\n", foo)
}
