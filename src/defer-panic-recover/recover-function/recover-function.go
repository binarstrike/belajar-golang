package main

import "fmt"

/*
Menurut ChatGPT:
Fungsi recover dalam bahasa pemrograman Go (Golang) adalah mekanisme yang digunakan untuk menghentikan panic
dan mengembalikan kendali program ke titik yang benar-benar dapat mengatasi situasi. Ini memungkinkan
program untuk melanjutkan eksekusi setelah terjadi panic, daripada menghentikannya sepenuhnya. Recover biasanya
digunakan dalam blok defer untuk menangani panic dengan lebih baik dan menghindari kegagalan total aplikasi.
*/
func handleError() {
	message := recover()
	if message != nil {
		logger(fmt.Sprintf("Terdapat error pada aplikasi: %v", message))
	}
}

func logger(message string) {
	fmt.Println("Log:", message)
}

func StartApplication(err bool) {
	// fungsi logger akan dijalankan ketika fungsi StartApplication selesai dijalankan
	// dan walaupaun fungsi panic dipanggil fungsi yang sebelumnya diberi keyword defer akan tetap dijalankan
	defer logger("Aplikasi selesai dijalankan")
	defer handleError()
	if err {
		// jika fungsi panic dipanggil maka akan mengakhiri eksekusi fungsi saat ini dan akan menjalankan fungsi
		// yang di defer. Mungkin simpel nya: jika panic ke trigger maka akan pergi ke atas dan mencari fungsi yang di defer
		// kemudian menjalankannya, untuk fungsi recover lebih baik dijalanakn di dalam fungsi yang di defer untuk
		// menghandle error pada panic. Fungsi recover akan mengambil data atau nilai yang ada pada parameter pada saat
		// fungsi panic dipanggil
		panic("Aplikasi error!!") // aplikasi akan langsung berhenti
		fmt.Println("Hi guys")    // baris kode ini tidak akan dijalankan
	}
	fmt.Println("Memulai menjalankan aplikasi")
}

func main() {
	StartApplication(false)

	fmt.Println("=====================")

	StartApplication(true)
	// karena error dari panic sudah di handle oleh recover maka aplikasi akan tetap berjalan
	// baris kode di bawah akan dijalankan
	fmt.Println("Hello World")
}
