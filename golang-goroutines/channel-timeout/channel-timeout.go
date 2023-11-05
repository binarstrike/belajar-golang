package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan string)

	go sendData(ch1)
	retrieveData(ch1)
}

func sendData(ch chan<- string) {
	randomizer := rand.NewSource(time.Now().Unix())

	for i := 0; true; i++ {
		randomTime := time.Duration(randomizer.Int63()%10+1) * time.Second
		ch <- fmt.Sprintf("data %d", i)
		fmt.Println("turu", randomTime)
		time.Sleep(randomTime)
	}
}

// label `loop` untuk block statement perulangan for
// label pada golang biasa digunakan untuk mengendalikan aliran eksekusi program
// dimana jika digunakan untuk perulangan dan didalam block perulangan tersebut terdapat kondisi dimana harus
// menghentikan perulangan dengan menggunakan label

func retrieveData(ch <-chan string) {

	// terdapat dua kondisi pada select..case dimana salah satunya menunggu data masuk/diterima dari channel
	// dan satunya adalah sebuah fungsi yang menngembalikan channel dari dengan waktu tunggu selama 5 detik.
	// jika tidak data dierima dari channel pada case pertama selama 5 detik maka akan menjalankan block kode
	// pada case kedua
loop:
	for {
		select {
		case data := <-ch:
			fmt.Printf("data diterima: %s\n", data)

			// case <-time.After(time.Second * 5):
		case <-bangunWoy(5):
			fmt.Println("timeout, turu lebih dari 5 detik")

			// bila tidak menggunakan label, keyword break disini hanya untuk menghentikan/membatalkan proses menjalankan
			// kode pada block switch..case jadi perulangan for tetap berjalan
			break loop
		}
	}
}

func bangunWoy(s time.Duration) <-chan bool {
	t := make(chan bool)
	go func(chTime chan bool) {
		time.Sleep(time.Second * s)
		chTime <- true
	}(t)
	return t
}

/*
contoh output dari terminal:
turu 2s
data diterima: data 0
turu 4s
data diterima: data 1
turu 4s
data diterima: data 2
turu 9s
data diterima: data 3
timeout, turu lebih dari 5 detik
*/
