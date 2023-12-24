package main

import (
	"fmt"
	"time"
)

// pada golang untuk berbagi atau bertukar data pada goroutine bisa mneggunakan `channel`.
// channel digunakan untuk proses bertukar data antara menjadi pengirirm atau penerima dimana data tersebut
// akan menumpuk pada channel dan menunggu untuk diterima atau nilai dari data tersebut diambil dan disimpan
// oleh variabel tertentu yang tipe data nya sesuai dengan pada saat deklarasi channel

// ketika sebuah data pada channel sudah diterima oleh penerima maka data tersebut akan dihilangkan oleh channel
// dan akan diisi dengan data selanjutnya sesuai antrian

// pada saat menjalankan aplikasi, go runtime akan membuat satu goroutine utama untuk menjalankan aplikasi
// atau fungsi main
func main() {
	// ada dua jenis channel yaitu un-buffered channel dan buffered channel
	// pada saat deklarasi un-buffered channel tidak perlu ditentukan ukuran atau panjang buffer nya
	channel1 := make(chan string) // make(chan string, 0)
	var str1 string

	go func() {
		time.Sleep(time.Second * 3)
		channel1 <- "Ucup" // pengirim atau yang mengisi data ke channel
	}()

	// proses akan terhenti sampai str1 mendapatkan atau terisi data dari channel
	str1 = <-channel1 // penerima atau variabel yang akan mengambil data dari channel
	// fmt.Println("Halo", <-channel1) // channel juga bisa digunakan langsung pada parameter dari sebuah fungsi
	fmt.Println("Halo", str1)

	concurrrentFib(10)
}

func concurrrentFib(n int) {
	chInts := make(chan int)

	go func() {
		for i := n; i > 0; i-- {
			fibonacciChan(i, chInts)
		}
	}()

	for i := n; i > 0; i-- {
		fmt.Println(<-chInts)
	}
}

func fibonacciChan(n int, c chan int) {
	c <- fibonacci(n)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
