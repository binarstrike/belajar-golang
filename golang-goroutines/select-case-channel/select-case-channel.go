package main

import (
	"fmt"
)

// sama seperti switch..case select..case digunakan khusus untuk kondisi dimana channel yang terlebih dahulu menerima
// dari pengirim data pada goroutine lain maka block kode pada case/kondisi tersebut akan dijalankan

func main() {
	var (
		myNums []int = []int{1, 2, 3, 23, 12, 6, 60, 32, 8, 9, 10, 20, 21, 40}
		ch1          = make(chan int)
		ch2          = make(chan float64)
	)

	go max(myNums, ch1)
	go average(myNums, ch2)

	// jika salah satu case/kondisi terpenuhi maka akan mengabaikan case lainnya dan akan menjalankan block kode
	// pada case/kondisi dimana data lebih dulu sampai/selesai dikirim dan diterima
	for i := 0; i < 2; i++ {
		select {
		case maxNum := <-ch1:
			fmt.Printf("Max: %d\n", maxNum)

		case avg := <-ch2:
			fmt.Printf("Average: %3.f\n", avg)
		}
	}
}

func average(nums []int, ch chan float64) {
	var sums int = 0

	for _, v := range nums {
		sums += v
	}

	ch <- float64(sums) / float64(len(nums))
}

func max(nums []int, ch chan int) {
	maxNum := nums[0]

	for _, v := range nums {
		if maxNum < v {
			maxNum = v
		}
	}
	ch <- maxNum
}
