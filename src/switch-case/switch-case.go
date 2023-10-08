package main

import "fmt"

// golang juga memiliki fitur switch expression yaitu bentuk lain dari percabangan selain if-else
func main() {
	printMe := fmt.Println
	const nilai = 83

	switch true {
	case nilai >= 90:
		printMe("Predikat: A")
	case nilai >= 85:
		printMe("Predikat: B")
	case nilai >= 80:
		printMe("Predikat: C")
	case nilai >= 75:
		printMe("Predikat: D")
	case nilai >= 60, nilai <= 70: // bisa diberi lebih dari satu kondisi
		printMe("Predikat: E")
	default: // bila semua kondisi tidak terpenuhi maka akan menjalanken kode pada blok default
		printMe("Predikat: F")
	}

	const name = "Binar"

	switch name {
	case "Binar":
		printMe("Halo Binar")
	case "Ucup":
		printMe("Halo Ucup")
	case "Otong":
		printMe("Halo Otong")
	}

	// switch expression juga bisa menggunakan short statement
	const batasPanjangNama = 4

	switch nameLength := len(name); nameLength > batasPanjangNama {
	case true:
		printMe("Panjang nama melebihi batas")
	case false:
		printMe("Ok")
	}

	// percabangan dengan switch expression juga bisa bisa dilakukan tanpa kondisi jadi lebih mirip if-else
	// jika kondisi pada salah satu case terpenuhi atau true maka akan menjalankan kode pada pada blok case tersebut
	const nilaiAkhir int = 87
	switch {
	case nilaiAkhir >= 90:
		printMe("Predikat: A")
	case nilaiAkhir >= 85:
		printMe("Predikat: B")
	case nilaiAkhir >= 80:
		printMe("Predikat: C")
	case nilaiAkhir >= 75:
		printMe("Predikat: D")
	case nilaiAkhir >= 70, nilaiAkhir < 70: // bisa diberi lebih dari satu kondisi
		printMe("Predikat: E")
	}
}
