package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// artikel penting: https://reintech.io/blog/introduction-to-golang-sync-atomic-package

/*
sumber: https://www.baeldung.com/java-atomic-variables.

Why count ++ is not atomic?
This is because of the simple increment operation (counter++), which may look like an atomic operation,
but in fact is a combination of three operations: obtaining the value, incrementing, and writing the updated value back.
If two threads try to get and update the value at the same time, it may result in lost updates.
*/

/*
agar error berupa terjadi nya race condition pada program yang dijalankan dapat terlihat dan keluar pada layar terminal,
tambahkan flag `-race` pada saat menjalankan kode program
go run -race concurrent-atomic-operation.go
*/

func main() {
	var (
		num1 atomic.Int32
		num2 int32
		wg   sync.WaitGroup
	)

	justIncrement(&num2, &wg)
	atomicIncrement(&num1, &wg)
	wg.Wait()

	fmt.Println(num1.Load()) // 1000
	fmt.Println(num2)        // angka tidak pasti namun mendekati 1000
}

func justIncrement(num *int32, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// akan terjadi race condition dimana pada satu atau lebih proses goroutine secara bersamaan mengkases satu data yang sama
		// sehingga menjadi kacau
		go func() {
			*num++
			wg.Done()
		}()
	}
}

func atomicIncrement(num *atomic.Int32, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			num.Add(1)
			wg.Done()
		}()
	}
}
