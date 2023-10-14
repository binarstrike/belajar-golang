package main

import "fmt"

/*
Menurut ChatGPT:
Empty interface dalam bahasa pemrograman Go (Golang) adalah jenis antarmuka khusus yang tidak memiliki metode.
Dalam Go, antarmuka didefinisikan oleh sekumpulan metode, tetapi empty interface tidak memiliki metode apa pun.
Hal ini membuat empty interface sangat fleksibel karena dapat digunakan untuk menyimpan nilai dari berbagai
tipe data yang berbeda.
*/

// tipe data any adalah alias untuk inteface{}
func AnyFunc(n int) any {
	switch n {
	case 1:
		return "Otong"
	case 2:
		return 100
	case 3:
		return true
	default:
		return nil
	}
}

func main() {
	var foo any
	var bar interface{} = 1

	foo = 42
	bar = "Ucup"
	fmt.Println(foo)
	fmt.Println(bar)

	fmt.Println()

	var1 := AnyFunc(1)
	var2 := AnyFunc(2)
	// var2 int := AnyFunc(2) // baris kode ini akan error karena memberi tipe data secara eksplisit sedangkan
	// tipe data yang dikembalikan oleh fungsi AnyFunc dapat berupa tipe data apa saja
	var3 := AnyFunc(3)
	var4 := AnyFunc(4)

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
}
