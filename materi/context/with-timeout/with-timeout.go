package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go turu(ctxTimeout, ch)

	// jika waktu atau jeda dari menjalankan fungsi doSomething lebih dari waktu timeout maka kondisi pertama akan terpenuhi
	// dimana context akan segera mengirim data ke chanel pada fungsi .Done() setelah waktu yang diberikan berakhir lebih dulu
	select {
	case <-ctxTimeout.Done():
		fmt.Println("Prit... pritt... timeout!")
	case str := <-ch:
		fmt.Println(str)
	}

}

func turu(ctx context.Context, ch chan string) {
	fmt.Println("turu for 3 second")
	time.Sleep(time.Second * 3)
	fmt.Println("bangun")
	ch <- "time to work"
}
