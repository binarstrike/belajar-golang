package main

import "fmt"

func main() {
	// ada beberapa cara untuk membuat variabel pada go
	/**
	  menggunakan keyword var
	  var nama_variabel <tipe_data>

	  atau bisa langsung dengan memberikan nilai ke variabel tersebut dan compiler akan menegenali secara otomatis
	  tipe data dari variabel tersebut.
	  var nama_variabel = 1000

	  atau bahkan bisa tanpa harus menggunakan keyword var dengan menggunakan `:=`
	  nama_variabel := "Foo"
	  sama saja dengan
	  var nama_variabel = "Foo"
	  var nama_variabel string = "Foo"
	*/

	var foo_str string
	foo_str = "Foo"
	fmt.Println(foo_str)

	var bar_str = "Bar"
	fmt.Println(bar_str)

	hello_str := "Hello World"
	fmt.Println(hello_str)

	var (
		firstName = "Binar"
		lastName  = "Nugroho"
	)
	fmt.Println(firstName, lastName)
}
