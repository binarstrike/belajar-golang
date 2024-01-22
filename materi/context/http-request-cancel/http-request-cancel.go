package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	go runServer()
	time.Sleep(time.Millisecond * 200)

	clientRequest()
}

func clientRequest() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)

	// client akan membatalkan permintaan ketika sudah melebihi batas waktu yang ditentukan
	client := &http.Client{Timeout: time.Millisecond * 1000}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, _ := io.ReadAll(resp.Body)
	fmt.Println("response:", string(data))
}

func runServer() {
	fmt.Println("server: server is running")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// go func() {
		// 	time.Sleep(time.Millisecond * 100)
		// 	fmt.Fprint(w, "Hello World")
		// 	fmt.Println("server: sending response to client")
		// }()

		select {
		case <-time.After(time.Millisecond * 2000):
			fmt.Println("server: sending response to client")
			w.Header().Add("Content-Type", "text/plain")
			fmt.Fprint(w, "Hello World")
		case <-ctx.Done():
			fmt.Println("server: request canceled from client")
		}

		// <-ctx.Done()
		// fmt.Println("server: request canceled from client")
	}))
}
