package main

import "fmt"

func main() {
	const nilaiAkhir uint8 = 86
	const absensi uint8 = 70
	const isLulus bool = nilaiAkhir >= 85 && absensi >= 70
	// && -> jika kedua nilai yang dibandingkan hasilnya true atau benar maka true
	// || -> jika salah satu nilai yang dibandingkan true maka true

	fmt.Println(isLulus)
	if isLulus {
		fmt.Println("Kamu lulus")
		return
	}
	fmt.Println("Kamu tidak lulus")
}
