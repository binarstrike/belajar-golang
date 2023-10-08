package main

import "github.com/fatih/color"

// sama seperti bahasa pemrograman lain golang juga memiliki fitur percabangan
func main() {
	printMe := color.New(color.FgHiCyan)
	const nilaiTest uint8 = 80
	if nilaiTest >= 80 {
		printMe.Println("Kamu lulus")
	} else {
		printMe.Println("Kamu tidak lulus")
	}

	printMe.Printf("\n%s\n", "====================================")
	if nilaiTest >= 90 {
		printMe.Println("Predikat A")
	} else if nilaiTest >= 85 {
		printMe.Println("Predikat B")
	} else if nilaiTest >= 80 {
		printMe.Println("Predikat C")
	} else if nilaiTest >= 75 {
		printMe.Println("Predikat D")
	} else {
		printMe.Println("Predikat E")
	}

	printMe.Printf("\n%s\n", "====================================")
	// golang mempunyai fitur if short statement atau ketika ingin melakukan pengecekan pada kondisi tertentu
	// bisa ditambahkan sebuah statement sebelum pengecekan kondisi
	const batasPanjangData = 5
	data := [...]int{1, 2, 3, 4, 5, 6}

	if dataLength := len(data); dataLength > batasPanjangData {
		printMe.Println("Panjang data melebihi batas")
		printMe.Printf("\ndata: %#v\n", data)
	}
}
