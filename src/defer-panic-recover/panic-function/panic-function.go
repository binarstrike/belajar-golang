package main

import "fmt"

/*
Menurut ChatGPT:
Panic function dalam bahasa pemrograman Go (Golang) adalah sebuah mekanisme yang digunakan untuk menghentikan
eksekusi program secara mendadak ketika terjadi kondisi yang tidak terduga atau fatal.
Biasanya, panic digunakan untuk mengatasi situasi yang tidak dapat ditangani oleh program,
seperti ketika terjadi kesalahan runtime yang serius. Ketika panic terjadi, program akan segera berhenti,
dan runtime akan mencetak pesan panic bersama dengan stack trace untuk membantu dalam debugging.
Panic digunakan sebagai tanda bahwa terjadi kesalahan yang sangat serius,
dan biasanya tidak digunakan dalam kasus penanganan kesalahan yang lebih umum.
*/
func logger(message string) {
	fmt.Println("Log:", message)
}

func StartApplication(err bool) {
	// fungsi logger akan dijalankan ketika fungsi StartApplication selesai dijalankan
	// dan walaupaun fungsi panic dipanggil fungsi yang sebelumnya diberi keyword defer akan tetap dijalankan
	defer logger("Aplikasi selesai dijalankan")
	if err {
		logger("Terdapat error pada aplikasi")
		panic("Aplikasi error!!") // aplikasi akan langsung berhenti
		fmt.Println("Hi guys")    // baris kode ini tidak akan dijalankan
	}
	fmt.Println("Memulai menjalankan aplikasi")
}

func main() {
	StartApplication(false)
	fmt.Println("=====================")
	StartApplication(true)
	fmt.Println("Hello World")
}
