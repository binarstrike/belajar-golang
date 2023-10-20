package main

import (
	"errors"
	"fmt"
	"time"
)

// error interface biasanya digunakan untuk return value pada sebuah fungsi jika fungsi tersebut harus mengembalikan
// sebuah data error jika terjadi kesalahan saat menjalankan fungsi tersebut.
// karena merupakan sebuah interface maka jika tidak terjadi error bisa mengembalikan nil
func bagi[Type int | float64](x, y Type) (Type, error) {
	if y == 0 {
		return 0, errors.New("tidak bisa dibagi dengan 0")
	}
	return x / y, nil
}

type MyError struct {
	When time.Time
	What string
}

// implementasi method Error() sesuai kontrak dari interface error
func (e MyError) Error() string {
	return fmt.Sprintf("%v - %s", e.When, e.What)
}

func run(i int) (string, error) {
	switch i {
	case 1:
		return "", &MyError{
			When: time.Now().Local(),
			What: "Something wrong",
		}
	default:
		return "Ok", nil
	}
}

func main() {
	cobaBagi1, err := bagi(10, 0)
	cobaBagi2, _ := bagi(105.7, 12)

	if err != nil {
		// fmt.Println(err)
		fmt.Println(err.Error())
	}

	fmt.Println(cobaBagi1)
	fmt.Println(cobaBagi2)

	str, errrrrrr := run(1)
	if errrrrrr != nil {
		// fungsi atatu method fmt.Println() akan secara otomatis mencari dan menjalankan method Error()
		// jika terdapat pada nilai dari type yang mengimplementasikan method Error() dari interface error
		fmt.Println(errrrrrr)
	} else {
		fmt.Println(str)
	}
}
