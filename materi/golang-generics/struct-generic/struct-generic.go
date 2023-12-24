package main

import "fmt"

// Comparable adalah sebuah interface khusus yang memuat type data yang dapat dibandingkan dengan operasi perbandingan
// seperti equal `==`, not equal `!=`, greater than `>`, less than `<` juga operasi perbandingan lainnya

type Person[T comparable] struct {
	Name string
	Age  T
}

func main() {
	udin := Person[int]{
		Name: "Udin",
		Age:  18,
	}
	ucup := Person[float64]{
		Name: "ucup",
		Age:  19.7,
	}

	fmt.Printf("udin: %#v\n", udin)
	fmt.Printf("udin: %#v\n", ucup)
}
