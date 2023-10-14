package main

import (
	"fmt"

	p "github.com/inancgumus/prettyslice"
)

/*
nil adalah nilai nol atau tidak ada dalam bahasa Go. Dalam konteks Go, nil sering digunakan untuk mewakili
ketiadaan nilai pada tipe data tertentu, seperti pointer, slice, map, channel, dan interface.
*/
func isDahMakan(makan string) any {
	if makan == "sudah" {
		return true
	}
	return nil
}

func NewUserSlice(name string) []string {
	if name == "" {
		return nil
	}
	return append(make([]string, 0), name)
}

func main() {
	makan := "sudah"
	dahMakan := isDahMakan(makan)

	if dahMakan == nil {
		fmt.Println("Belum makan")
	}
	fmt.Println("Sudah makan")

	fmt.Println()

	user1 := NewUserSlice("Dodi")

	if user1 == nil {
		fmt.Println("Data kosong")
		return
	}
	p.Show("Data", user1)
}
