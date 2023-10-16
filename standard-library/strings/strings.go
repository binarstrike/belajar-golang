package main

import (
	"fmt"
	s "strings"
)

var print = fmt.Println
var f = fmt.Sprintf

func main() {
	const str string = "--Binar Nugroho-"
	const str2 string = "Budi Budi Budi Ucup Otong Budi"

	strTrim := s.Trim(str, "-")          // menghilangkan karakater pada parameter cutset pada string
	strContains := s.Contains(str, "-")  // mengecek jika string terdapat karakter "-"
	strToLowerCase := s.ToLower(strTrim) // merubah string menjadi lowercase (huruf kecil) semua
	strToUpperCase := s.ToUpper(strTrim) // merubah string menjadi uppercase (huruf besar/kapital) semua
	strSplit := s.Split(strTrim, " ")    // membagi atau memisahkan string berdasarkan karaketer pada parameter sep
	strRepaceAll := s.ReplaceAll(str2, "Budi", "Dika")
	// mengubah semua string dari karaketer/string pada parameter old ke string baru pada parameter new

	print(f("strings.Trim: %s", strTrim))
	print(f("string.Contains: %v", strContains))
	print(f("strings.ToLower: %s", strToLowerCase))
	print(f("strings.ToUpper: %s", strToUpperCase))
	print(f("strings.Split: %#v", strSplit[:]))
	print(f("strings.ReplaceAll: %s", strRepaceAll))
}
