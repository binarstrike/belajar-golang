package main

import (
	"fmt"

	// fungsi init pada package pkginit akan dijalankan secara otomatis ketika di import
	pkginit "github.com/binarstrike/belajar-golang/pkg/pkg-init"

	// jika hanya ingin import sebuah package dan hanya menjalankan fungsi init nya saja
	// tambahkan `_` pada syntax import nya. Ini disebut Blank identifier
	_ "github.com/binarstrike/belajar-golang/pkg/pkg-blank-init"
)

func main() {
	fmt.Println(pkginit.GetFoo())
}
