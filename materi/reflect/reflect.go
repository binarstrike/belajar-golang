package main

import (
	"fmt"
	r "reflect"
	"strconv"

	"github.com/binarstrike/belajar-golang/pkg/person"
)

func main() {
	p1 := person.Person{FirstName: "Ucup", LastName: "Surucup", Age: 0}
	p1Value := r.ValueOf(p1)
	p1Type := r.TypeOf(p1)

	// fungsi .NumField()  mengembalikan nilai angka berupa jumlah dari berapa banyak field yang terdapat pada struct
	fmt.Println(p1Type.NumField())

	// mengambil nama field pertama pada struct Person
	fmt.Println(p1Type.Field(0).Name)
	// mengambil nama field ketiga pada struct
	fmt.Println(p1Type.Field(2).Name)

	// mengambil nilai dari struct tag `json` di field pertama pada struct Person
	fmt.Println(p1Type.Field(0).Tag.Get("json"))
	// mengambil nilai dari struct tag `json` di field kedua pada struct Person
	fmt.Println(p1Type.Field(1).Tag.Get("json"))

	// fungsi strconv.ParseBool() digunakan untuk mengonversi string dengan
	// nilai boolean yang terdapat pada sebuah string "true" atau "false"
	isLastNameFieldOptional, _ := strconv.ParseBool(p1Type.Field(1).Tag.Get("required"))
	fmt.Printf("type: %T\nvalue: %v\n\n", isLastNameFieldOptional, isLastNameFieldOptional)

	for i := 0; i < p1Type.NumField(); i++ {
		// mengambil nama json key pada struct tag `json` dan nilai dari setiap field pada struct
		fmt.Printf("%s: %v\n", p1Type.Field(i).Tag.Get("json"), p1Value.Field(i).Interface())
	}

	if isP1Valid, err := IsValid(p1); err != nil {
		fmt.Println(err)
		fmt.Println(isP1Valid)
	} else {
		fmt.Println("p1 is valid")
	}
}

func IsValid(s any) (bool, error) {
	value := r.ValueOf(s)
	type_ := r.TypeOf(s)

	for i := 0; i < type_.NumField(); i++ {
		if a := type_.Field(i); a.Tag.Get("required") == "true" {
			switch t := value.Field(i).Interface(); t.(type) {
			case string:
				if t != "" {
					break // break disini akan mengakhiri pengencekan kondisi pada switch...case dan melanjutkan perulangan
				}
				return false, fmt.Errorf("field %s is not a valid value: %v", a.Name, t)
			case int, uint, float64, complex64:
				if t != 0 {
					break
				}
				return false, fmt.Errorf("field %s is not a valid value: %v", a.Name, t)
			}
		}
	}
	return true, nil
}
