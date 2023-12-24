package main

import "fmt"

// type assertion digunakan untuk mengubah nilai dari tipe interface kosong interface{} menjadi tipe aslinya
// type assertion mengembalikan dua nilai yaitu nilai dari hasil konversi dan kondisi berhasil atau tidak operasi
// type assertion tersebut dengan tipe bool

// any adalah type alias untuk inteface{}
func Foo(y int) any {
	switch y {
	case 1:
		return "Hello"
	case 2:
		return 7
	case 3:
		return true
	default:
		return nil
	}
}

func main() {
	var x interface{} = 10

	if value, isOk := x.(int); isOk {
		fmt.Println("x:", value)
	} else {
		fmt.Println("type assertion gagal")
	}

	if val, isOk := Foo(1).(int); isOk {
		fmt.Println(val)
	} else {
		// bila type assertion gagal maka val akan memiliki nilai 0 sebagai default value dari int
		fmt.Println(val, "type assertion gagal")
	}

	if val, isOk := Foo(1).(bool); isOk {
		fmt.Println(val)
	} else {
		fmt.Println(val, "type assertion gagal") // val = false
	}

	if val, isOk := Foo(3).(bool); isOk {
		fmt.Println(val)
	} else {
		fmt.Println(val, "type assertion gagal")
	}

	// bisa juga menggunakan switch case
	switch value := Foo(1).(type) {
	case string:
		fmt.Println("tipe data string", value)
	case int:
		fmt.Println("tipe data int", value)
	case bool:
		fmt.Println("tipe data boolean", value)
	default:
		fmt.Println("tipe data tidak diketahui")
	}
}
