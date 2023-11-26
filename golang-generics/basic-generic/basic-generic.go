package main

// Generic merupakan salah satu fitur pada golang yang digunakan untuk membuat fungsi, struktur data dan interface agar dapat
// menerima tipe data secara dinamis.

type UserIDInt int
type UserIDFloat float64

// AddUser fungsi ini dapat menerima tipe data int dan float64
func AddUser[T int | float64](userId T) T {
	return userId
}

// AddUser2 penggunaan karakter atau operator tilde `~` pada type parameter digunakan agar fungsi generic dapat menerima
// tipe data alias dari tipe data sebenarnya
func AddUser2[T ~int | ~float64](userId T) T {
	return userId
}

func main() {
	_ = AddUser(2)
	_ = AddUser(3.14)

	var userIdIntA UserIDInt = 20
	var userIdFloatB UserIDFloat = 2.189
	// _ = AddUser(id) // baris kode ini akan error karena type parameter tidak dapat menerima tipe data selain int dan float64
	userA := AddUser2(userIdIntA)
	_ = userA
	userB := AddUser2(userIdFloatB)
	_ = userB
}
