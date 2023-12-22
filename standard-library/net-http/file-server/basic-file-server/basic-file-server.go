package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()

	root := http.FileServer(http.Dir("."))
	static := http.FileServer(http.Dir("static"))

	handler.Handle("/", http.StripPrefix("/", root))
	handler.Handle("/static/", http.StripPrefix("/static/", static))

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
