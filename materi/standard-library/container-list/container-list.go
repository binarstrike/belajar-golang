package main

import (
	"container/list"

	a "github.com/binarstrike/belajar-golang/pkg/alias"
)

func main() {
	data := list.New()

	data.PushBack("Binar")   // Binar
	data.PushBack("Nugroho") // Binar Nugroho
	data.PushBack("Dika")    // Binar Nugroho Dika
	data.PushFront("Ucup")   // Ucup Binar Nugroho Dika
	data.PushFront(100)      // 100 Ucup Binar Nugroho Dika

	// data sebenarnya yang terdapat pada list data adalah
	// <nil> 100 Ucup Binar Nugroho Dika <nil>
	// nilai nil tersebut terus ada pada urutan pertama list dan akhir list
	// jenis struktur data ini disebut `double linked list`

	a.Print(a.F("data length: %d", data.Len()))

	for el := data.Front(); el != nil; el = el.Next() {
		a.Print(a.F("%v", el.Value))
	}

	// berdasarkan isi data sekarang data.Front().Value = Binar, maka data.Front().Prev() = nil
	// karena data.Front().Prev() = nil maka jika mengambil nilai nil dari data.Front().Prev().Value
	// akan terjadi panic.
	a.Print(func(el *list.Element) any {
		if el != nil {
			return el.Value
		}
		return nil
	}(data.Front().Prev()))
}
