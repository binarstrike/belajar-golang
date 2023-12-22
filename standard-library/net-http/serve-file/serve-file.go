package main

import (
	"fmt"
	"net/http"
)

func main() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
