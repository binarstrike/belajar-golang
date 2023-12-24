package main

import "fmt"

// selain bisa mengembalikan nilai lebih dari satu fungsi pada golang juga bisa menamai variabel dari nilai yang akan
// dikembalikan
func fullName() (firstName string, middleName string, lastName string) {
	firstName = "Ucup"
	middleName = "Albert"
	lastName = "Santoso"

	// bisa diubah urutan variabel dari nilai yang akan dikembalikan
	// return firstName, lastName, middleName

	// atau juga bisa langsung menggunakan keyword return tanpa menambah variabel yang akan dikembalikan
	// urutan akan menyesuaikan pada urutan di deklarasi fungsi
	return
}

func main() {
	first, middle, last := fullName()
	fmt.Printf("Namaku adalah: %s %s %s\n", first, middle, last)
}
