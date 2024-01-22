package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("time is finish")
	case <-ctx.Done():
		fmt.Println("process canceled")
	}

}
