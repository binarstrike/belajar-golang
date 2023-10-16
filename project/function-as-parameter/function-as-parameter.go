package main

import p "github.com/inancgumus/prettyslice"

// Function type declaration
// digunakan untuk mempermudah pembuatan fungsi yang menerima fungsi sebagai parameter nya
// ini akan berguna jika terdapat lebih dari satu parameter pada fungsi
// juga dapat dipakai ulang pada pembuatan fungsi lain (reusable)
type FilterType[Type any] func(Type) bool

// parameter fungsi pada golang juga bisa berupa fungsi
// sama seperti callback atau higher order function pada javascript

// ada beberapa cara deklarasi fungsi jika ingin menggunakan fungsi sebagai parameter
// func Filter[Type any](array []Type, filter func(Type) bool) []Type {

func Filter[Type any](array []Type, filter FilterType[Type]) []Type {
	filterSlice := make([]Type, 0)

	for _, value := range array {
		if filter(value) {
			filterSlice = append(filterSlice, value)
		}
	}
	return filterSlice
}

func FilterAngkaGenap(num int) bool {
	return num%2 == 0
}

func FilterWarnaBiru(color string) bool {
	return color == "Biru"
}

func main() {
	// jika menggunakan `...` pada deklarasi array maka ini adalah array
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15}

	// bila nilai yang akan di input pada parameter fungsi Filter adalah array biasa maka diubah dulu menjadi slice
	// dengan menambahkan `:` di antara kurung siku
	p.Show("Angka ganjil", Filter[int](numbers[:], func(num int) bool {
		return num%2 != 0
	}))

	// juga dapat langsung menerima parameter berupa fungsi yang sudah di deklarasi dan dibuat sebelumnya
	p.Show("Angka genap", Filter(numbers[:], FilterAngkaGenap))

	// ini slice
	colors := []string{"Biru", "Merah", "Biru", "Hijau", "Merah", "Kuning", "Merah", "Biru", "Biru", "Orange"}
	p.Show("Ditemukan warna merah", Filter[string](colors, func(warna string) bool {
		return warna == "Merah"
	}))

	p.Show("Ditemukan warna biru", Filter[string](colors, FilterWarnaBiru))
}
