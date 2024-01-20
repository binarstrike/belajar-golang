package main

import (
	"fmt"
	"runtime"
	"time"
)

// artikel penting: https://dasarpemrogramangolang.novalagung.com/A-buffered-channel.html

// dengan menggunakan buffered channel data yang dikirim ke sebuah channel dapat ditahan di memori tanpa harus ada
// penerima yang siap menerima data sebagai antrian sesuai dengan ukuran buffer pada channel saat deklarasi channel

func main() {
	runtime.GOMAXPROCS(2)

	// nilai buffered channel dimulai dari 0 jadi jika nilai ukuran yang dimasukan 3 maka jumlah buffer maksimal = 4
	c1 := make(chan string, 3)
	tnow := time.Now()

	go func() {
		// delay untuk melihat kalau terjadi blocking di pengiriman data pada channel
		time.Sleep(time.Second * 1)
		for {
			fmt.Println("data diterima <-:", <-c1, time.Duration(time.Since(tnow).Milliseconds()))
		}
	}()

	// data akan menumpuk pada antrian dengan maksimal data yang bisa ditumpuk 3 dan jika belum ada goroutine lain
	// yang siap menerima data dari channel c1 maka akan terjadi blocking pada perulangan ke 3 dan akan menunggu hingga
	// ada goroutine lain yang akan menerima data.
	for i := 1; i <= 10; i++ {
		// time.Sleep(time.Millisecond * 100)
		c1 <- fmt.Sprint("data ", i)
		fmt.Println("data terkirim ->: data", i, time.Duration(time.Since(tnow).Milliseconds()))
	}

	time.Sleep(time.Second * 3)

}

/*
Karena proses eksekusi nya yang sangat cepat jadi ketika ada gouroutine lain yang mengambil data dari buffer atau antrian
data baru akan langsung dimasukan karena ada antrian yang kosong dan akan terus berlanjut hingga perulangan selesai dan
itu terjadi pada saat perulangan pengiriman data ke 5 dan menerima data ke 5 hingga seterusnya.

contoh output:
data terkirim ->: data 1 110ns
data terkirim ->: data 2 218ns
data terkirim ->: data 3 327ns
data terkirim ->: data 4 2.008µs
data diterima <-: data 1 2.008µs
data diterima <-: data 2 2.008µs
data diterima <-: data 3 2.008µs
data diterima <-: data 4 2.008µs
data terkirim ->: data 5 2.115µs
data diterima <-: data 5 2.115µs
data terkirim ->: data 6 2.225µs
data diterima <-: data 6 2.225µs
data terkirim ->: data 7 2.334µs
data diterima <-: data 7 2.334µs
data terkirim ->: data 8 2.444µs
data diterima <-: data 8 2.444µs
data terkirim ->: data 9 2.552µs
data diterima <-: data 9 2.553µs
data terkirim ->: data 10 2.661µs
data diterima <-: data 10 2.662µs
*/
