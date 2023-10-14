package main

import "fmt"

// - Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
// - Sebuah interface berisikan definisi-definisi method
// - Biasanya interface digunakan sebagai kontrak

type IAnimal interface {
	GetDescription() (jenisHewan string, namaHewan string, jumlahKaki int)
}

type Animal struct {
	NamaHewan, JenisHewan string
	JumlahKaki            int
}

func DisplayAnimalDescription(animal IAnimal) {
	jenisHewan, namaHewan, jumlahKaki := animal.GetDescription()
	fmt.Printf("\nJenis: %s\nNama hewan: %s\nJumlah kaki: %d\n", jenisHewan, namaHewan, jumlahKaki)
}

func (animal Animal) GetDescription() (jenisHewan string, namaHewan string, jumlahKaki int) {
	return animal.JenisHewan, animal.NamaHewan, animal.JumlahKaki
}

// bayangkan saja seperti kontrak bahwa jika ada sebuah fungsi yang memiliki parameter dengan tipe sebuah
// type interface maka setiap variabel atau instance dari struct yang akan menjadi parameter dari
// fungsi tersebut harus membuat atau memiliki fungsi sesuai fungsi-fungsi pada interface.
// mungkin mirip dengan keyword implements yang digunakan saat deklarasi class pada bahasa pemrograman OOP
// seperti TypeScript atau Java
func main() {
	kelinci := Animal{
		NamaHewan:  "Kelinci",
		JenisHewan: "Mamalia",
		JumlahKaki: 4,
	}

	ayam := Animal{
		NamaHewan:  "Ayam",
		JenisHewan: "Unggas",
		JumlahKaki: 2,
	}

	DisplayAnimalDescription(kelinci)
	DisplayAnimalDescription(ayam)
}
