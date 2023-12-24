package main

import (
	"fmt"
	"sync"
)

// link penting: https://dasarpemrogramangolang.novalagung.com/A-mutex.html

// sync.Mutex
// sync.Mutex digunakan untuk mencegah terjadinya "race condition" sebuah kondisi dimana sebuah data
// diakases secara bersamaan baik itu mengambil data atau menulis nilai baru pada data tersebut.
// dengan menggunakan sync.Mutex data ini bisa dikunci/lock sementara dan kemudian akan dibuka kembali
// ketika proses/kode perintah yang mengakses data ini telah selesai mengakses nya.
// dan ketika ada proses yang berjalan sebagai contoh berupa goroutine mencoba mengakses data ini secara bersamaan
// dengan goroutine lainnya maka mereka akan secara bergantian untuk mengakses data tersebut dan akan menumpuk seperti
// antrian jika terjadi banyak sekali proses goroutine yang mencoba akses data tersebut.

//	type Counter struct {
//		value int
//	}
type Counter struct {
	sync.Mutex
	value int
}

//	func (c *Counter) Add() {
//		c.value++
//	}
func (c *Counter) Add() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	var wg sync.WaitGroup
	var counter_1 Counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter_1.Add()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter_1.Value())
}
