package main

import (
	"fmt"

	p "github.com/inancgumus/prettyslice"
)

/*
Menurut ChatGPT:
Closure adalah konsep dalam pemrograman di mana sebuah fungsi memiliki akses ke variabel-variabel yang
berada di luar lingkup fungsi tersebut, bahkan setelah lingkup fungsi tersebut selesai dieksekusi.
Dengan kata lain, fungsi yang merupakan closure dapat "mengingat" dan "mengakses" variabel-variabel
yang ada dalam lingkup parent atau fungsi yang mengelilinginya. Hal ini memungkinkan untuk mempertahankan dan
memanipulasi data dalam lingkup yang lebih luas, sering digunakan dalam pemrograman untuk membuat fungsi yang
lebih fleksibel dan berguna dalam berbagai konteks.
*/
func main() {
	names := make([]string, 0, 10)

	addName := func(name string) {
		names = append(names, name)
	}

	outer := func() func(int) int {
		counter := 0

		inner := func(n int) int {
			counter += n
			return counter
		}

		return inner
	}

	counter1 := outer()
	counter2 := outer()
	fmt.Println(counter1(1)) // +1
	fmt.Println(counter1(3)) // 1 + 3
	fmt.Println(counter1(5)) // 4 + 5
	fmt.Println()
	fmt.Println(counter2(5)) // +5
	fmt.Println(counter2(5)) // 5 + 5

	addName("Ucup")
	addName("Otong")
	p.Show("", names)
	addName("Dika")
	addName("Ani")
	addName("Joko")
	p.Show("", names)
}
