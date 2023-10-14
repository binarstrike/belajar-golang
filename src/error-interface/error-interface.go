package main

import (
	"errors"
	"fmt"
)

// error interface biasanya digunakan untuk return value pada sebuah fungsi jika fungsi tersebut harus mengembalikan
// sebuah data error jika terjadi kesalahan saat menjalankan fungsi tersebut.
// karena merupakan sebuah interface maka jika tidak terjadi error bisa mengembalikan nil
func bagi[Type int | float32](x, y Type) (Type, error) {
	if y == 0 {
		return 0, errors.New("tidak bisa dibagi dengan 0")
	}
	return x / y, nil
}

func main() {
	cobaBagi1, err := bagi(10, 0)
	cobaBagi2, _ := bagi[float32](105.7, 12)

	if err != nil {
		// fmt.Println(err)
		fmt.Println(err.Error())
	}

	fmt.Println(cobaBagi1)
	fmt.Println(cobaBagi2)
}
