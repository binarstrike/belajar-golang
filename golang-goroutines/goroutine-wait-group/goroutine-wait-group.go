package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup
// sync.WaitGroup adalah sebuah struct atau fitur pada builtin library pada golang yang berguna untuk
// menunggu sebuah goroutine atau banyak goroutine yang dijalankan dalam satu waktu sebelum menjalankan perintah lainnya
// dengan menambahkan sebuah angka pada fungsi .Add() yang akan digunakan counter atau penghitung berapa banyak
// proses goroutine yang dijalankan.

// memanggil fungsi .Wait() akan menghentikan sementara proses menjalankan kode dan akan menunggu hingga counter menjadi 0
// dan semua proses goroutine selesai dijalankan.

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		d := fmt.Sprintf("data %d", i)

		wg.Add(1)
		// pastikan sebelum menjalankan fungsi goroutine harus memnaggil fungsi wg.Add(1) untuk menambah counter
		// pada sync.WaitGroup bahwa proses goroutine yang berjalan bertambah 1
		go printMe(&wg, d)
	}

	wg.Wait()
	fmt.Printf("==================================================\n\n")

	chName := make(chan string)
	for _, name := range []string{"Ucup", "Dika", "Ani", "Udin", "Budi", "Otong"} {
		wg.Add(2)
		go sayHi(&wg, chName)
		go sendName(&wg, name, chName)
	}
	wg.Wait()
}

func printMe(wg *sync.WaitGroup, s string) {
	// pastikan untuk memanggil fungsi wg.Done() pada akhir dari fungsi goroutine yang sedang berjalan untuk mengurangi
	// counter pada sync.WaitGroup sebagai tanda bahwa ada proses goroutine yang sudah selesai dijalankan.

	// jika wg.Done() tidak dipanggil maka akan terjadi panic jika counter pada sync.WaitGroup bukan sama dengan 0
	// fatal error: all goroutines are asleep - deadlock!

	// penggunaan defer disini agar memanggil fungsi wg.Done() saat semua statement/perintah pada block kode fungsi
	// yang dipanggil dengan goroutine selesai.
	defer wg.Done()
	fmt.Println(s)
}

func sayHi(wg *sync.WaitGroup, chName <-chan string) {
	defer wg.Done()
	fmt.Println("Hi", <-chName)
}

func sendName(wg *sync.WaitGroup, name string, ch chan<- string) {
	defer wg.Done()
	ch <- name
}
