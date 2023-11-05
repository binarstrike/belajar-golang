package main

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/binarstrike/belajar-golang/pkg/person"
)

// yang penting nama key pada struct tag json sama dengan nama key pada json string yang akan di unmarshal atau
// diubah menjadi data yang bisa diakses runtime dan memiliki type
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}

func main() {
	p1 := person.Person{FirstName: "Otong", LastName: "Santoso", Age: 18}
	p2 := map[string]string{
		"name": "Ucup",
		"age":  "17",
	}
	// 	encoding/marshalling sebuah object dari struct atau map ke json string.
	// Saat melakukan encoding atau marshal dengan method/fungsi json.Marshal() fungsi akan mengembalikan dua nilai
	// sebuah data dengan tipe []byte berupa json yang masih dalam bentuk `byte code` yang sudah di encoding dari data berupa object struct atau map dan juga error.
	// Error akan menjadi nilai nil jika tidak terjadi error pada saat melakukan proses encoding.
	p1Json, err := json.Marshal(p1)
	p2Json, _ := json.Marshal(p2)
	if err != nil {
		fmt.Println("terdapat error ketika encoding object p1")
	}
	// karena bentuk dan tipe kembalian dari funsgi json.Marshal masih berbentuk byte maka harus dilakukan konversi tipe data
	// ke string
	fmt.Printf("p1: %s\n", string(p1Json))
	fmt.Printf("p1: %s\n", string(p2Json))

	jsonStr_1 := `
  {
    "name": "Izumi",
    "age": 16,
    "hobby": "Drawing"
  }`
	jsonStr_2 := `
  {
    "name": "Rikka",
    "age": 17,
    "hobby": "Drawing"
  }`
	jsonStr_3 := fmt.Sprintf("[%s,%s]", jsonStr_1, jsonStr_2)

	var userIzumi User

	// json.Unmarshal([]byte(json1String), &userIzumi)
	// json.Unmarshal() dapat mengembalikan error jika terjadi kesalahan pada proses unmarshal json string
	// parameter kedua dari json.Unmarshal() adalah pointer ke sebuah variabel yang menyimpan perubahan hasil
	// perubahan dari json dalam bentuk string ke data yang memiliki tipe data sesuai dengan bentuk data dan tipe
	// pada properti-properti string json
	err_ := json.Unmarshal([]byte(jsonStr_1), &userIzumi)

	if err_ != nil {
		fmt.Println("terdapat error ketika mengubah json string menjadi data dengan tipe User")
	}
	fmt.Printf("userIzumi: %#v\n", userIzumi)

	var userRikka any
	json.Unmarshal([]byte(jsonStr_2), &userRikka)
	userRikkaMap := userRikka.(map[string]any)

	fmt.Printf("userRikkaMap age type: %T\n", userRikkaMap["age"])
	fmt.Printf("userRikkaMap: %#v\n", userRikkaMap)

	var userArrayIzumiRikka []User

	json.Unmarshal([]byte(jsonStr_3), &userArrayIzumiRikka)
	fmt.Printf("userArrayIzumiRikka: %#v\n", userArrayIzumiRikka)
	// userArrayIzumiRikka: []main.User{main.User{Name:"Izumi", Age:16, Hobby:"Drawing"}, main.User{Name:"Rikka", Age:17, Hobby:"Drawing"}}
}
