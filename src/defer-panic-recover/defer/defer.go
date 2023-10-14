package main

import "fmt"

// defer adalah sebuah keyword yang digunakan untuk memeanggil fungsi ketika fungsi utama selesai dijalankan
func logger(message string) {
	fmt.Println("Log:", message)
}

func StartApplication() {
	// fungsi logger akan dijalankan ketika fungsi StartApplication selesai dijalankan
	defer logger("Aplikasi selesai dijalankan")
	fmt.Println("Memulai menjalankan aplikasi")
}

func main() {
	StartApplication()
}
