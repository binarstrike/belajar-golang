package main

import "fmt"

/**
  syntax:
  var nama_variaabel <tipe data>
  docs: https://pkg.go.dev/builtin#int
*/

// signed integer
var var1_int int
var var1_int8 int8   // -128 - 127
var var1_int16 int16 // -32768 - 32767
var var1_int32 int32 // -2147483648 - 2147483647
var var1_int64 int64 // -9223372036854775808 - 9223372036854775807

// unsigned integer
var var2_uint8 uint8   // 0 - 255
var var2_uint16 uint16 // 0 - 65535
var var2_uint32 uint32 // 0 - 4294967295
var var2_uint64 uint64 // 0 - 18446744073709551615

// floating point
var var3_float32 float32       // 1.18 * 10^-38 - 3.4 * 10^38
var var3_float64 float64       // 2.23 * 10^-308 - 1.80x10^308
var var3_complex64 complex64   // https://pkg.go.dev/builtin#complex64
var var3_complex128 complex128 // https://pkg.go.dev/builtin#complex128

// boolean
var var4_boolean bool // true / false

// string
var var5_string string // sebuah tipe data berupa kumpulan karakter yang di apit dalam tanda petik dua atau kutip

func main() {
	// signed integer
	var1_int = 1
	var1_int8 = 127
	var1_int16 = -32768
	var1_int32 = 2147483647
	var1_int64 = 9223372036854775807
	fmt.Println("signed integer: ", var1_int, var1_int8, var1_int16, var1_int32, var1_int64)

	// unsigned integer
	var2_uint8 = 255
	var2_uint16 = 65534
	var2_uint32 = 4294967294
	var2_uint64 = 18446744073709551613
	fmt.Println("unsigned integer: ", var2_uint8, var2_uint16, var2_uint32, var2_uint64)

	// floating point
	var3_float32 = 12.3
	var3_float64 = 1212312312.12121212
	var3_complex64 = 1.78758758475878499787587687858
	var3_complex128 = 1.7736378647647588758665785746
	fmt.Println("floating point: ", var3_float32, var3_float64, var3_complex64, var3_complex128)

	// boolean
	var4_boolean = true
	fmt.Println("boolean true: ", var4_boolean)
	var4_boolean = false
	fmt.Println("boolean false: ", var4_boolean)

	// string
	var5_string = "Hello World"
	fmt.Println("string: ", var5_string)
	var5_string = "Kururin..."
	fmt.Println("string: ", var5_string)
}
