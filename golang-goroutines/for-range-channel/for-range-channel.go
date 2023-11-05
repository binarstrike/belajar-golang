package main

import "fmt"

// artikel penting: https://dasarpemrogramangolang.novalagung.com/A-channel-range-close.html

// channel juga bisa digunakan pada perulangan for dengan tambahan keyword range dimana ini akan memicu untuk menjalankan
// block kode pada perulangan ketika ada data yang masuk atau diterima pada channel.

// disini juga akan memanfaatkan pembatasan level akses pada suatu channel yang dinamakan channel direction.
// Dengan menggunakan fitur ini penngembang bisa membatasi akses pada channel seperti hanya dapat mengirim data saja atau
// hanya dapat menerima data saja dan juga bisa kedua nya bisa menerima dan mengirim.
// Contoh:
/*
  ch <- chan string: hanya dapat menerima data saja
  ch chan <- string: hanya dapat mengirim data saja
  ch chan string: bisa mengirim dan menerima data
*/
func main() {
	chName := make(chan string)

	go sendName(chName)
	printName(chName)
}

func sendName(ch chan<- string) {
	names := []string{"Ucup", "Dika", "Andi", "Budi", "Rikka", "Yuuki", "Chitanda"}

	for _, name := range names {
		// fmt.Printf("Send name: %s\n", name)
		ch <- name
	}
	close(ch)
}

func printName(ch <-chan string) {
	for name := range ch {
		fmt.Printf("Name : %s\n", name)
	}
}
