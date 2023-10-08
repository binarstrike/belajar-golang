package main

import "fmt"

// fungsi
func sapaUcup() {
	fmt.Println("Hi Ucup")
}

func iniSiOtong() string {
	return "Otong"
}

func sapaAku(nama string) {
	fmt.Printf("\nHalo %s\n", nama)
}

func main() {
	for i := 0; i < 5; i++ {
		sapaUcup()
	}

	// memanggil fungsi dengan nilai kembalian
	fmt.Printf("\nHalo %s\n", iniSiOtong())

	// memanggil funsi dengan parameter
	sapaAku("Udin")
}
