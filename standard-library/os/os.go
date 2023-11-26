package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Args nilai pada argument saat binary hasil compile dijalankan
	args := os.Args

	// contoh: os.exe binar nugroho
	// atau
	// go run os.go binar nugroho
	// Args: []string{"/path/to/os.exe", "binar", "nugroho"}
	fmt.Printf("Args: %#v\n", args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Printf("Hostname: %s\n", hostname)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
	openFile()
}

func openFile() {
	fileFoo, err := os.Open("./foo.json")

	if err != nil {
		fmt.Println("error when trying to open file")
		log.Fatal(err)
	}
	// setelah melakukan seleai melakukan perubahan baik hanya membaca atau menulis file jangan lupa untuk menutupnya.
	// bisa dengan keyword defer agar fungsi ini dijalankan saat fungsi utama selesai dijalankan dan pastikan untuk
	// menaruhnya dibawah baris kode pengecekan error dari variabel error pada fungsi os.Open()
	defer fileFoo.Close()

	// library os juga dipakai untuk membuka file system yang ada pada host
	// terdapat dua cara untuk menyimpan hasil file yang dibaca yaitu dengan membuat variabel array dengan panjang yang ditentukan
	// atau membuat slice dengan fungsi make() yang bisa langsung bisa diakses oleh fungsi .Read() tanpa mengubahnya menjadi slice
	readFooJson := [32]byte{}
	// readFooJson := make([]byte, 32)

	// untuk membaca file bisa menggunakan fungsi/method .Read() pada instance dari fungsi os.Open()
	// fungsi ini mengembalikan dua nilai jumlah total karakter yang terbaca dan error.
	// variabel n disini adalah variabel yang menyimpan panjang atau total jumlah character yang terbaca pada file
	if n, err := fileFoo.Read(readFooJson[:]); err != nil {
		fmt.Println("error when trying to read foo.json file")
		log.Fatal(err)
	} else {
		fmt.Printf("\ntotal karakter: %d\n", n)
	}

	fmt.Println(readFooJson)
	fmt.Printf("foo.json content: \n%s\n", string(readFooJson[:]))
	fmt.Println("readFooJson cap:", cap(readFooJson))
	fmt.Println("readFooJson len:", len(readFooJson))
}
