package main

import "fmt"

// sama seperti bahasa lain ada beberapa macam operator untuk melakukan operasi perbandingan pada go
func main() {
	const num1 uint8 = 100
	const num2 uint8 = 70
	const num1GreaterThanNum2 bool = num1 > num2
	const num1GreaterThanEqualNum2 bool = num1 >= num2
	const num1LastThanNum2 bool = num1 < num2
	const numLastThanEqualNum2 bool = num1 <= num2
	const num1EqualNum2 bool = num1 == num2
	const num1NotEqualNum2 bool = num1 != num2

	// %t untuk format string boolean
	// \n adalah karakter newline atau `enter`
	fmt.Printf("%t\n%t\n%t\n%t\n%t\n%t\n", num1GreaterThanNum2, num1GreaterThanEqualNum2, num1LastThanNum2,
		numLastThanEqualNum2, num1EqualNum2, num1NotEqualNum2)
}
