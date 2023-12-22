package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

// salah satu cara untuk menyimpan file-file seperti file html, css dan gambar untuk keperluan konten yang akan ditampilkan
// pada client dari web server yaitu menyimpannya langsung pada file binary hasil output/compile dengan menggunakan
// fitur embed pada golang

//go:embed static
var staticEmbedDir embed.FS

//go:embed pages
var pagesEmbedDir embed.FS

func main() {
	mux := http.NewServeMux()

	pagesSubFS, err := fs.Sub(pagesEmbedDir, "pages")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	staticSubFS, err := fs.Sub(staticEmbedDir, "static")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// https://stackoverflow.com/a/73227566/17546933
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(pagesSubFS))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticSubFS))))

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
