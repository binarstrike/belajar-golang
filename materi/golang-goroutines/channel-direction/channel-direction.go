package main

import "fmt"

// Channel direction

// sebuah channel dapat dibatasi akses yaitu hanya bisa mengirim atau menerima saja dan keduanya bisa mengirim dan menerima data
// dengan pembatasan membuat kode jauh lebih aman karena channel hanya diberikan akses yang diperlukan saja
/*
  ch <- chan string: hanya dapat menerima data saja
  ch chan <- string: hanya dapat mengirim data saja
  ch chan string: bisa mengirim dan menerima data
*/
func main() {
	ch1 := make(chan string)
	go sendOnly(ch1)
	receiveOnly(ch1)
}

func sendOnly(ch chan<- string) {
	ch <- "Ucup"
}

func receiveOnly(ch <-chan string) {
	fmt.Println("Hai", <-ch)
}
