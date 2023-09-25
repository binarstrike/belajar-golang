package main

import "fmt"

func main() {
	fooNumber := 10000000000000
	barNumber := fmt.Sprint(fooNumber) // konversi angka ke string
	fmt.Println(barNumber)

	name := "Binar"
	b := name[0] // variabel ini akan memiliki nilai sebuah karakter `B` dalam bentuk byte yaitu 66
	fmt.Println(b)
	bString := string(b) // perlu dilakukan konversi dari byte ke string dengan method atau fungsi `string`
	fmt.Println(bString)

	var angka1_int32 int32 = 200
	var angka1_int64 int64 = int64(angka1_int32)
	var angka1_int8 int8 = int8(angka1_int64) // -56
	// nilai pada variabel `angka1_int8` akan terjadi `Integer overflow` @see https://en.wikipedia.org/wiki/Integer_overflow
	// simpel nya ketika sebuah variabel dengan tipe data yang memiliki panjang yang kecil seperti int8 yaitu
	// dari -128 sampai 127 dan diberi data yang memiliki data diluar jangkauan tersebut maka akan nilai akan terus mengulang
	// ketika sudah mencapai panjang maksimum kemudian ke nilai yang terkecil lagi dan seterusnya

	fmt.Println(fmt.Sprintf("\n\n%d\n%d\n%d", angka1_int32, angka1_int64, angka1_int8))
}
