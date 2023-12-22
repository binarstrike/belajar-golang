package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// parameter kedua pada http.HandlerFunc yaitu *http.Request memuat informasi
	// tentang request http dari client seperti http method yang digunakan, url path yang diakses dan lain-lain
	var handler http.HandlerFunc = func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(
			writer,
			"URL Path: %s\nMethod: %s\nQuery parameter: %s",
			r.URL.Path,
			r.Method,
			r.URL.Query().Encode(),
		)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// jika mengkases url berikut melalui browser
// http://localhost:8080/hello/ucup?foo=bar
// maka keluaran atau data yang akan dikirim dari server:
/*
URL Path: /hello/ucup
Method: GET
Query parameter: foo=bar
*/
