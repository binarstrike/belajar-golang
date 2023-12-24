package main

import (
	"fmt"

	// klik ctrl + left mouse pada nama package dibawah untuk mengarahkan ke source file
	a "github.com/binarstrike/belajar-golang/pkg/animal"
)

// fmt.Stringer adalah interface yang hanya memiliki method String() string
// method ini akan dicari dan pertama kali dijalankan jika method diimplementasikan pada sebuah nilai
// dan dipassing pada sebuah fungsi dari yang mengimplememntasikan method ini contohnya fmt.Print
func main() {
	a1 := a.Animal{
		Name: "Kelinci",
		Type: a.Mammal,
	}

	a1_name, a1_type := a.GetNameAndType(a1)

	fmt.Printf("Animal name: %s\nAnimal type: %s\n", a1_name, a1_type)
}
