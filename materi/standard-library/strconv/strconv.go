package main

import (
	"fmt"
	"log"
	s "strconv"
)

var (
	print = fmt.Println
	f     = fmt.Sprintf
)

func main() {
	const strNum = "32"
	const strNumNegative = "-64"
	const luckyNumber = 13

	strToInt(strNum)
	strToInt("ok")
	intToStr(luckyNumber)
	parseInt(strNumNegative)
	parseInt("hello")
	formatBool(true)
	formatBool(false)
}

func handleError() {
	message := recover()

	if message != nil {
		log.Println(message)
	}
}

// mengubah tipe string yang terdapat angka pada nilainya ke tipe int
func strToInt(str string) {
	result, err := s.Atoi(str)

	defer handleError()

	if err == nil {
		print(f("strconv.Atoi: (%T) %#v", result, result))
	} else {
		panic(err.Error())
	}
}

// mengubah angka dengan tipe int ke string
func intToStr(num int) {
	result := s.Itoa(num)

	print(f("strconv.Itoa: (%T) %#v", result, result))
}

// secara default .ParseInt mengembalikan tipe int64 walaupun parameter bitSize diisi nilai selain 64
// tetap bisa dilakukan konversi dengan hasil yang valid
func parseInt(str string) {
	result, err := s.ParseInt(str, 10, 8)

	defer handleError()

	result_ := int8(result)
	if err == nil {
		print(f("strconv.ParseInt: (%T) %#v", result_, result_))
	} else {
		panic(err.Error())
	}
}

// mengubah nilai boolean true/false dengan tipe bool ke string
func formatBool(b bool) {
	result := s.FormatBool(b)

	print(f("strconv.FormatBool: (%T) %#v", result, result))
}
