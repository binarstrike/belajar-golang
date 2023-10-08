package main

import "fmt"

// variadic function adalah function yang bisa menerima parameter atau argument dengan jumlah yang
// bervariasi dan dipisah dengan tanda koma `,`. Parameter ini hanya dapat ditaruh di akhir urutan
// parameter pada deklarasi fungsi parameter ini disimpan pada variabel sebagai slice.
// Pada Javascript fitur ini di sebut `rest parameter`
func plus(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(plus(10, 10, 10, 10))

	numbers := []int{10, 20, 30, 40, 50, 60, 70}
	// fungsi ini juga dapat menerima parameter berupa slice dengan menambah `...` pada parameter
	total := plus(numbers...)
	fmt.Println(total)
}
