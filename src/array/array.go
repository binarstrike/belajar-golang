package main

import (
	"fmt"

	pretty "github.com/inancgumus/prettyslice"
)

func filterInt(arr []int, f func(int) bool) []int {
	myArray := make([]int, 0)

	for _, el := range arr {
		if f(el) {
			myArray = append(myArray, el)
		}
	}
	return myArray
}

func main() {
	var names [2]string
	names[0] = "Binar"
	names[1] = "Nugroho"

	fmt.Printf("Nama ku adalah %s %s\n", names[0], names[1])

	luckNumber := []int{7, 13, 16, 3, 100} // deklarasi arary secara langsung atau inline dengan nilai nya
	pretty.Show("My luck number: ", luckNumber)

	myNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	pretty.Show("Angka genap: ", filterInt(myNums, func(num int) bool { return num%2 == 0 }))
	pretty.Show("Angka ganjil: ", filterInt(myNums, func(num int) bool { return num%2 != 0 }))
}
