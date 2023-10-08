package main

import (
	"fmt"
	"sort"
)

// keyword `break` dan `continue` digunakan untuk menghentikan perulangan perbedaan nya adalah
// break digunakan untuk menghentikan seluruh perulanagan sedaangkan continue menghentikan
// proses perulangan yang sedang berjalan dan melanjutkan ke perulangan selanjutnya
func main() {
	printMe := fmt.Println
	f := fmt.Sprintf

	AbsenMurid := make(map[string]string)
	AbsenMurid["Ucup"] = "hadir"
	AbsenMurid["Udin"] = "alpa"
	AbsenMurid["Dika"] = "hadir"
	AbsenMurid["Otong"] = "hadir"
	AbsenMurid["Ani"] = "alpa"

	for key, value := range AbsenMurid {
		if value == "alpa" {
			printMe(f("%s tidak hadir", key))
			continue
		}
		printMe(f("%s hadir", key))
	}

	printMe()
	AntiMerah := map[int]string{
		1: "Biru",
		2: "Kuning",
		3: "Hijau",
		4: "Merah",
		5: "Biru",
	}
	// -- baris kode untuk mengurutkan key atau kunci pada map berupa angka
	numbers := make([]int, 0, len(AntiMerah))
	for k := range AntiMerah {
		numbers = append(numbers, k)
	}
	sort.Ints(numbers)
	// --
	for _, number := range numbers {
		if AntiMerah[number] == "Merah" {
			printMe(f("ditemukan warna Merah pada nomor %d", number))
			// saat ditemukan string Merah di salah satu data pada map maka perulangan akan berhenti
			break
		}
		printMe(f("%d: %s", number, AntiMerah[number]))
	}
}
