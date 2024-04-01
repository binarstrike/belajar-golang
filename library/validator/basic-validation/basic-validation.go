package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type User struct {
	Name  string `validate:"required,min=5,max=16"`
	Age   int    `validate:"required,min=17,max=25"`
	Email string `validate:"required,email"`
}

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	ucup := &User{
		Name:  "Ucup", // error, minimal jumlah karakter  5
		Age:   16,     // error, minimal nilai angka untuk field Age 17
		Email: "ucup24@xyz.com",
	}

	anton := &User{
		Name:  "Anton",
		Age:   22,
		Email: "anton$example.com", // error, alamat email tidak valid
	}

	otong := &User{
		Name:  "Otong",
		Age:   20,
		Email: "otong.blah@example.com",
	}

	validateStruct(ucup)
	validateStruct(anton)
	validateStruct(otong)
}

func validateStruct(s interface{}) {
	errs := validate.Struct(s)
	if errs != nil {
		fmt.Printf("%T\n\n", errs)

		if v, ok := errs.(validator.ValidationErrors); ok {
			for _, err := range v {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}

}
