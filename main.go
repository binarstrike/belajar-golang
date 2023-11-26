package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []string{
		"ok",   // 0
		"yes",  // 1
		"foo",  // 2
		"bar",  // 3
		"vlah", // 4
		"bruh", // 5
		"nice", // 6
	}

	a = slices.Delete[[]string](a, 4, 5)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", a[2])
}
