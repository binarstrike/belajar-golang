package main

import (
	color "github.com/fatih/color"
	pretty "github.com/inancgumus/prettyslice"
)

/*
Detail Tipe Data Slice
  - Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
  - Pointer adalah penunjuk data pertama di array para slice
  - Length adalah panjang dari slice, dan
  - Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

Membuat Slice dari Array
  - array[start:end] Membuat slice dari array dimulai index low sampai index sebelum high
  - array[start:] Membuat slide dari array dimulai index low sampai index akhir di array
  - array[:end] Membuat slice dari array dimulai index 0 sampai index sebelum high
  - array[:] Membuat slice dari array dimulai index 0 sampai index akhir di array
*/
func main() {
	yellowBg := color.New(color.BgYellow)
	printMe := color.New(color.Bold, color.FgHiGreen)
	days := [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// len: 2
	// cap: 2
	// Sabtu, Minggu
	hariLibur := days[5:] // mengambil data dari indeks 5 hingga indeks terakhir

	pretty.Show("Hari Libur:", hariLibur)

	// saat melakukan slice pada sebuah array maka array tersebut saling terhubung atau saling reference dari
	// variabel yang terdapat slice tersebut. Jadi jika terjadi perubahan data pada variabel yang terdapat array yang
	// di slice maka array asal atau aslinya akan ikut berubah juga.
	hariLibur[0] = "Senin"
	hariLibur[1] = "Senin"
	pretty.Show("Mana hari minggu:", days[:])

	// append()
	// fungsi append digunakan untuk menambah data baru pada indeks terkahir sebuah array namun jika
	// data pada array sudah penuh maka akan membuat array baru dengan data yang ditambahkan tersebut
	yellowBg.Println("=================================================")
	colors := [7]string{"Red", "Blue", "Cyan", "Yellow", "Green", "Aquamarine", "Magenta"}
	pretty.Show("colors", colors[:])

	colorSlice1 := colors[5:] // Aquamarine, Magenta | dari indeks 5 hingga akhir
	pretty.Show("colorSlice1 colors[5:]", colorSlice1)

	colorSlice1Append := append(colorSlice1, "Pink")
	pretty.Show("colorSlice1 append(colorSlice1, \"Pink\")", colorSlice1Append) // Cyan, Magenta, Pink
	pretty.Show("colors", colors[:])
	// Red, Blue, Cyan, Yellow, Green, Aquamarine, Magenta
	// colorSlice1Append menjadi array baru yang tidak ter reference pada array asal yaitu colors
	// karena tidak dapat dilakukan penambahan data pada array slice yang sudah mencapai indeks terkahir
	yellowBg.Println("=================================================")

	pretty.Show("colors", colors[:])
	colorSlice2 := colors[1:4] // Blue, Cyan, Yellow | dari indeks 1 hingga indeks 3
	pretty.Show("colorSlice2 colors[1:4]", colorSlice2)

	colorSlice2Append := append(colorSlice2, "Pink")
	pretty.Show("colorSlice2 append(colorSlice2, \"Pink\")", colorSlice2Append) // Blue, Cyan, Yellow, Pink
	pretty.Show("colors", colors[:])
	// Red, Blue, Cyan, Yellow, Pink, Aquamarine, Magenta
	// data pada array colors akan ikut terganti karena saat melakukan penambahan data pada array slice
	// colorSlice2 belum mencapai indeks terakhir jadi data akan ditambahakan dan menggantikan data
	// pada indeks 4 yaitu "Green"
	yellowBg.Println("=================================================")

	printMe.Printf("\n%d\n%s\n", cap(colors[3:]), colors[3:])
}
