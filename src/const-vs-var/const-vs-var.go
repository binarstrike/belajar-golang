package main

import "fmt"

/**
bila mendeklarasikan sebuah variabel dengan `const` maka variabel tersebut tidak akan bisa diubah nilai nya
namun jika menggunakan `var` nilai dari variabel tersebut dapat dibuah seterusnya.
*/

func main() {
	var fooStr string = "Foo"
	const barSecretNumber uint8 = 250

	fmt.Println(fooStr)
	fooStr = "Foo Bar" // mengubah nilai dari variabel `fooStr`
	fmt.Println(fooStr)

	fmt.Println(barSecretNumber)
	// barSecretNumber = 10; // ini akan error
}
