package main

import "fmt"

func main() {
	bar := map[string]uint{
		"A": 100,
		"B": 90,
		"C": 80,
		"D": 70,
	}

	// tipe data barKeys akan manjadi tipe data []string sesuai tipe ada pada map yang dimasukan pada parameter fungsi
	barKeys := getKeys(bar)

	fmt.Println(barKeys)
}

func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
