package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

// salah satu cara untuk menyimpan file-file seperti file html, css dan gambar untuk keperluan konten yang akan ditampilkan
// pada client dari web server yaitu menyimpannya langsung pada file binary hasil output/compile dengan menggunakan
// fitur embed pada golang

//go:embed static pages
var data embed.FS

func main() {
	mux := http.NewServeMux()

	pagesDir, _ := fs.Sub(data, "pages")
	staticDir, _ := fs.Sub(data, "static")

	// https://stackoverflow.com/a/73227566/17546933
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(pagesDir))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticDir))))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// }

	// if err := server.ListenAndServe(); err != nil {
	// 	fmt.Println(err)
	// }

}
