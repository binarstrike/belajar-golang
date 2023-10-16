package main

import (
	"fmt"
)

// golang hanya mempunyai satu bentuk perulangan yaitu for-loop
func main() {
	printMe := fmt.Println
	f := fmt.Sprintf
	count := 0

	for count <= 5 {
		printMe(f("Halo Ucup ke %d", count))
		count++
	}
	printMe()
	// perulangan pada golang juga mendukung short statement sama seperti percabangan
	for i := 1; i < 6; i++ {
		printMe(f("Halo Otong ke %d", i))
	}

	printMe()
	// memanfaatkan perulangan untuk iterasi data pada array
	names := [2]string{"Ucup", "Otong"}

	for i := 0; i < len(names); i++ {
		printMe(f("indeks ke %d: %s", i, names[i]))
	}

	printMe()
	// atau bisa juga menggunakan keyword `range`
	for _, name := range names {
		// variabel yang pertama sebelah kiri adalah indeks dari array dan yang ke dua adalah nilai dari array
		// jika tidak ingin menggunakan variabel indeks bisa diganti dengan underscore
		printMe(f("Halo %s", name))
	}

	printMe()
	// for range juga bisa digunakan pada map
	Students := make(map[int]string)
	Students[1] = "Ucup"
	Students[2] = "Otong"
	Students[3] = "Dika"
	Students[4] = "Udin"

	for key, value := range Students {
		printMe(f("Murid %d: %s", key, value))
	}
}
