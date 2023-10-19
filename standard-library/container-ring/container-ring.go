package main

import (
	"container/ring"
	"fmt"
)

/*
ring atau circular list
Menurut ChatGPT:
Ring atau circular list adalah jenis struktur data yang mirip dengan linked list,
tetapi dengan perbedaan bahwa elemen terakhir dalam list menghubungkan kembali ke elemen pertama,
membentuk loop. Ini membuatnya berguna dalam situasi di mana kita perlu mengakses elemen-elemen list
dalam urutan berulang, seperti dalam pengulangan.
*/

func main() {
	data := ring.New(5) // 0 <-> 1 <-> 2 <-> 3 <-> 4 <-> 0 <-> 1 <-> ...

	for i := 0; i < /* 5 */ data.Len(); i++ {
		data.Value = fmt.Sprintf("Croissant %d", i)
		data = data.Next()
		// saat perulangan sudah mencapai kondisi false (5 < 5) data kembali ke 0
	}

	// .Do() digunakan untuk iterasi semua nilai pada data
	data.Do(func(v any) {
		fmt.Println(v)
	})

	fmt.Println()

	// panjang data = 5
	// sekarang data berada pada posisi pertama 0
	// 0 -> 1 -> 2 pindah 2 langkah ke posisi 2 dimana pada data memiliki nilai "Croissant 2"
	data = data.Move(2)
	fmt.Println(data.Value)

	// 2 -> 3 -> 4 -> 0 -> 1 -> ... pindah 3 langkah ke posisi 0 "Croissant 0" dan akan terus berputar ketika berpindah posisi
	data = data.Move(3)
	fmt.Println(data.Value)

	// posisi data berada pada 0 dan kembali ke posisi sebelum 0 yaitu 4
	data = data.Prev()
	fmt.Println(data.Value)
}
