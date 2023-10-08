package main

import "fmt"

// fungsi pada golang juga bisa mengembalikan lebih dari satu nilai
// tipe data dari nilai yang dikembalikan juga harus disesuiakan dengan tipe dan jumlah nilai yang di kembalikan
func myName() (string, string) {
	return "Binar", "Nugroho"
}

func main() {
	firstName, lastName := myName()
	fmt.Printf("Namaku adalah %s %s\n", firstName, lastName)

	// jika ingin menghiraukan nilai kembalian dari memanggil fungsi bisa menggunakan tanda garis bawah atau undescore `_`
	namaDepan, _ := myName()
	fmt.Printf("\nHi %s\n", namaDepan)
}
