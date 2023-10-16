package cobainit

import "fmt"

var foo string

// fungsi init adalah fungsi spesial yang akan dieksekusi ketika package lain meng-import package ini
func init() {
	// fungsi akan dijalankan dan mengisi nilai pada variabel foo
	fmt.Println("Aku telah dipanggil!")
	foo = "Foo"
}

func GetFoo() string {
	return foo
}
