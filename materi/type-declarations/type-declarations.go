package main

import "fmt"

// type declarations adalah membuat tipe data baru dari tipe data yang sudah ada
// atau bisa disebut alias dari sebuah tipe data. Untuk membuat type alias menggunakan keyword `type`
// contoh syntax:
// type <nama alias> <tipe data>
// type NoKTP string

type Skor uint8
type SudahMakan bool

func main() {
	var skorBola Skor = 10
	fmt.Println(skorBola)

	var isUcupSudahMakan SudahMakan = true
	if isUcupSudahMakan {
		fmt.Println("Ucup sudah makan")
	} else {
		fmt.Println("Ucup belum makan")
	}
}
