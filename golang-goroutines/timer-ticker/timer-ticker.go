package main

import (
	"fmt"
	"time"
)

func main() {
	// ketika membuat instance baru dari time.NewTimer timer akan langsung berjalan
	// dengan memanfaatkan channel C pada struct time.Timer atau hasil nilai kembalian dari time.NewTimer
	// bisa untuk menghentikan proses eksekusi dengan menunggu data masuk ke channel C sesuai waktu yang tersisa dan yang sudah
	// ditentukan saat deklarasi timer
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("start timer")
	// time.Sleep(1 * time.Second)
	// fmt.Println("3 second left")
	// time.Sleep(2 * time.Second)
	// fmt.Println("1 second left")
	<-timer.C
	fmt.Println("finish timer")

	chBool := make(chan bool)

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("turu")
		chBool <- true
	})

	fmt.Println("start time.AfterFunc")
	<-chBool
	fmt.Println("finish time.AfterFunc")

	tickerTick()
}

// time.Ticker
// ticker biasanya digunakan untuk penjadwalan sebuah proses untuk dijalankan.
// struct time.Ticker memiliki field C dengan tipe channel time.Time dan akan menerima data dengan tipe time.Time
// sesuai jeda yang diatur saat deklarasi variabel dengan fungsi time.NewTicker secara terus-menerus
// dan akan berhenti ketika fungsi/method .Stop() dipanggil
func tickerTick() {
	done := make(chan bool)
	myTicker := time.NewTicker(300 * time.Millisecond)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
		fmt.Println("ticker berakhir")
	}()

	for {
		select {
		case <-done:
			myTicker.Stop()
			time.Sleep(time.Microsecond)
			return
		case t := <-myTicker.C:
			fmt.Println(t.Format(time.DateTime))
		}
	}
}
