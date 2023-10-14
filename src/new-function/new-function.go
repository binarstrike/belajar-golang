package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// fungsi new() digunakan untuk membuat pointer dengan nilai default dan tidak ada data atau nilai awal yang di assign
func main() {
	var num1 *int = new(int) // membuat pointer ke int
	var num2 **int = &num1   // pointer num2 ke pointer num1 ke pointer int
	var num3 ***int = &num2  // dengan operator ** satu maka itu alamat memori nya dan *** itu nilainya dan seterusnya....
	// pointer num3 ke pointer num2 ke pointer num1 ke pointer int

	fmt.Println("num1:", &num1) // 0x....
	fmt.Println("num2:", num2)  // 0x.... => &num1
	// jika hanya num3 tanpa operator * maka itu sama saja dengan &num2 sesuai pada saat deklarasi variabel
	fmt.Println("num3:", *num3) // 0x.... => num2
	fmt.Println()
	fmt.Println("num1:", num1)   // 0x....
	fmt.Println("num2:", *num2)  // 0x....
	fmt.Println("num3:", **num3) // 0x....

	fmt.Println()

	// untuk mengakses nilai dari pointer menggunakan operator atau karakter * didepan nama variabel pointer nya
	***num3 = 100
	fmt.Println("num1:", *num1)  // 100
	fmt.Println("num2:", **num2) // 100

	fmt.Println()

	structPointer()
}

func structPointer() {
	p1 := new(Person)
	p2 := &p1
	p3 := &p2

	fmt.Printf("p1: %#v\n", p1)

	p1.Name = "Ucup"
	fmt.Printf("p1: %#v\n", p1)

	(**p2).Age = 18
	fmt.Printf("p1: %#v\n", p1)

	**p3 = &Person{Name: "Otong", Age: 19}
	fmt.Printf("p1: %#v\n", p1)

}
