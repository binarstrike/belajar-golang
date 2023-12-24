package main

import (
	"fmt"
	"net/http"
	"os"
)

var handler http.HandlerFunc = func(w http.ResponseWriter, _ *http.Request) {
	// w adalah respon yang akan dikirim ke client atau web browser
	// fungsi fmt.Fprint digunakan untuk menulis data berupa string pada varaiabel parameter w
	// yang akan dikirim dan ditampilkan di sisi client.
	fmt.Fprintf(w, "%s %d\n", "Hello World", 1234)
	fmt.Fprint(w, "Halo Ucup")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
