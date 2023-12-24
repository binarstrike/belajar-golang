package main

import "fmt"

// function pada golang juga bisa disimpan ke variabel sebagai nilai dari variabel tersebut
func sapaAku(name string) {
	fmt.Println("Hi " + name)
}

func main() {
	sapa := sapaAku
	sapa("Otong")
}
