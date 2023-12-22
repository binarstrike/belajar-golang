package main

import (
	"fmt"
	"log"
	"net/http"
)

// allowOnlyGet middleware untuk check kalau hanya method GET yang diperbolehkan untuk melakukan request.
func allowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("only GET is allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// validateQueryParam middleware untuk validasi terhadap query parameter pada request url.
func validateQueryParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("name") == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("parameter name is empty"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// sayHello sebuah handler akhir atau controller yang akan mengembalikan data berupa string dengan isi sebuah sapaan
// terhadap nama yang diambil dari query parameter.
var sayHello http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	mux := http.NewServeMux()
	const port string = ":8080"

	// baris kode dibawah akan menerapkan middleware secara global atau untuk semua url path/endpoint akan melewati middleware ini terlebih dahulu
	// var handler http.Handler = mux
	// handler = validateQueryParam(handler)
	// handler = allowOnlyGet(handler)

	mux.Handle("/hello", allowOnlyGet(validateQueryParam(sayHello)))
	// mux.Handle("/hello", sayHello)

	log.Printf("server started at localhost%s\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		// if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}
