package main

import (
	"fmt"
	re "regexp"
)

func main() {
	defer handlePanic()

	// regex untuk mengecek jika terdapat kata yang memiliki 4 karakter dengan flag insensitive atau
	// mengabaikan besar dan kecil nya huruf
	r1 := re.MustCompile(`(?i)\b[a-z]{4}\b`)
	word1 := "Ucup Dika Anton Udin Ani Budi"

	// jika parameter n = -1 maka akan mencari atau mencocokan string tanpa batasan
	findStr := r1.FindAllString(word1, 5)
	fmt.Println(findStr)

	for _, name := range [2]string{"Otong", "Siti"} {
		if r1.MatchString(name) {
			fmt.Println("regex cocok dengan nama:", name)
			continue
		}
		fmt.Println("regex tidak cocok dengan nama:", name)
	}
}

func handlePanic() {
	msg := recover()

	if msg != nil {
		fmt.Println(msg)
	}
}
