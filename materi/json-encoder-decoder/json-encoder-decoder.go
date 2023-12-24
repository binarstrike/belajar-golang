package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
)

// Biasanya penggunaan dari package json/go-json biasa digunakan untuk menangani payload dengan format file json dari sebuah
// request/permintaan dan response dari sebuah API atau http server

type jsonForm struct {
	Hello string `json:"hello"`
	Foo   string `json:"foo"`
}

func main() {
	var cobaJson1 jsonForm = jsonForm{Hello: "World", Foo: "Bar"}
	file, _ := os.OpenFile(".\\file.json", os.O_CREATE, os.ModePerm)
	defer file.Close()

	jsonEnc := json.NewEncoder(file)

	cobaJson1.Hello = "Ucup"
	fmt.Printf("jsonHello: %#v\n", cobaJson1)

	if err := jsonEnc.Encode(cobaJson1); err != nil {
		log.Println(err)
	}
}
