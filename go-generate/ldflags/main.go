package main

import (
	"fmt"
)

// salah satu fitur yang berhubungan dengan go generate adalah penggunaan command flag
// -ldflags pada saat build atau menjalankan langsung kode program
// dimana ini dapat digunakan untuk mengubah nilai dari variabel yang terdapat pada
// kode program pada saat melakukan proses build tanpa merubahnya secara langsung
// dengan menulis nya di kode program.
// go run -ldflags "-X importPath.varName=value" main.go

// go run -ldflags "-X main.versionString=v2.0" ldflags.go

var versionString = "v1.0"

func main() {
	fmt.Println("version: ", versionString)
}
