package main

import (
	"fmt"
	"time"
)

// golang juga dapat menjalankan sebuah perintah atau fungsi secara asynchronous atau non-blocking
// dengan menggunakan keyword `go` pada saat memanggil fungsi, fungsi tersebut akan dieksekusi secara asynchronous
func main() {
	for i := 1; i < 6; i++ {
		go func(a int) {
			t := time.Second * (0 + time.Duration(a))
			time.Sleep(t)
			fmt.Println("Halo Ucup", t)
		}(i)

		// string "Halo Budi ...ms" akan ditampilkan pada terminal lebih dulu karena jeda nya dibawah 1 detik
		go func(a int) {
			t := time.Millisecond * (200 * time.Duration(a))
			time.Sleep(t)
			fmt.Println("Halo Budi", t)
		}(i)
	}
	fmt.Println("Done")

	// bila fungsi main selesai dijalankan maka fungsi yang berjalan pada goroutine akan dihentikan juga bersama
	// aplikasi yang dijalankan, untuk contoh kali ini sebelum proses berakhir diberi jeda selama 7 detik.
	time.Sleep(time.Second * 7)
}
