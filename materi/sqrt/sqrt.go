package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 1; i < 11; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Perulangan ke %d: z = %v\n", i, z)
	}
	return z
}

func main() {
	x := 25.0 // ✓25 = 5
	r1 := Sqrt(x)
	fmt.Printf("Hasil r1: %v\n", r1)
	fmt.Printf("Hasil akar kuadrat x dengan math.Sqrt(): %v\n", math.Sqrt(x))
}
