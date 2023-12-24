package main

import (
	"fmt"
	"net/http"
	"os"
)

// Agar dapat menangani banyak URL Path atau route sekaligus dapat menggunakan fungsi
// http.NewServeMux dimana fungsi ini akan mengembalikan struct http.ServerMux yang akan
// digunakan untuk menambahkan route atau url path baru serta fungsi callback yang akan
// menangani request tersebut.

func main() {
	var mux *http.ServeMux = http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, "Hello from /") // Hello from /
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi from ", request.URL.Path) // Hi from /hi
	})

	// perlu diperhatikan pada pattern parameter pertama, jika terdapat `/` pada
	// akhirannya maka jika melakukan request /hello/ dan /hello/any/path
	// akan mengembalikan data dari request untuk /hello namun tetap
	// memprioritaskan request untuk url path yang terdaftar pada fungsi handler.

	// http://localhost:8080/hello/ucup?foo=bar
	// hasil keluaran atau data yang akan dikirim dari server ketika mengkases url tersebut:
	//    Hello from /hello/ucup
	//    query parameter: foo=bar
	mux.HandleFunc("/hello/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(
			writer,
			"Hello from %s\nquery parameter: %s",
			request.URL.Path,
			request.URL.Query().Encode(),
		)
	})

	// ketika mengkases endpint atau url path web server akan mengutamakan
	// path yang lebih panjang dan terdaftar pada fungsi handler, namun jika
	// path tersebut tidak ditemukan atau terdaftar pada fungsi handler
	// maka url path akan mundur 1 langkah dan membuat request pada path tersebut
	// jika mengkases /hello/ucup maka yang ditampilkan adalah request untuk path
	// /hello karena /hello/ucup tidak terdaftar di fungsi handler.
	mux.HandleFunc("/hello/budi", func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, "Hello Budi") // Hello Budi
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
