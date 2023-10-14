package main

import (
	"fmt"
)

// fungsi rekursif adalah sebuah fungsi yang memenggil fungsi itu sendiri di dalam blok programnya
func FaktorialLoop(num int) int {
	result := 1
	for a := num; a > 0; a-- {
		result *= a
	}
	return result
}

func FaktorialRekursif(num int) int {
	if num == 1 {
		return 1
	}
	return num * FaktorialRekursif(num-1)
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(FaktorialLoop(5))
	fmt.Println(FaktorialRekursif(5))
	fmt.Println(Fibonacci(6))
}
