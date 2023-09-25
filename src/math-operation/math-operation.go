package main

import "fmt"

// sama seperti bahasa pemrograman lain go juga memakai operator yang sama untuk melakukan operasi matematika seperti
// +   -    *    /   dan modulo atau hasil bagi   %

func main() {
	const numInt1 int = 20
	const numInt2 int = 15
	const numFloat1 = 30.127
	const numFloat2 = 10.23

	const tambah = numInt1 + numInt2
	const kurang = numInt1 - numInt2
	const kali = numInt1 * numInt2
	const bagi float32 = numFloat1 / numFloat2
	const modulo = numInt1 % numInt2

	fmt.Printf("%d\n%d\n%d\n%.2f\n%d\n", tambah, kurang, kali, bagi, modulo)

	fmt.Println("================")
	fmt.Println("Augmented Assignment")
	// Augmented Assigment
	// go dapat melakukan operasi matematika dengan syntax yang lebih singkat jika operasi tersebut
	// dilakukan dengan variabel yang langsung ingin diubah nilai nya dengan melakukan operasi matematika

	num3 := 100
	num3 += 10 // ini sama saja dengan num3 = num3 + 10
	fmt.Println(num3)
	num3 -= 10
	fmt.Println(num3)
	num3 *= 10
	fmt.Println(num3)
	num3 /= 10
	fmt.Println(num3)
	num3 %= 23
	fmt.Println(num3)

	fmt.Println("=================")
	fmt.Println("Unary Operator")
	// Unary Operator
	// sama seperti bahasa pemrograman lainnya golang juga memiliki unary operator seperti `increment`, `decrement` dan NOT operator
	myNum1 := 25
	myNum1++            // increment, sama saja dengan blahNum = blahNum + 1
	fmt.Println(myNum1) // 26
	myNum1--
	fmt.Println(myNum1) // 25

	isMarried := false
	fmt.Println(!isMarried) // true, NOT atau kebalikan dari nilai variabel `isMarried` yaitu false menjadi true

	// nilai negatif dan positif
	var negatifNum int8 = -18
	var positifNum = +32 // secara default semua angka pada go adalah positif tanpa harus menyertakan tanda atau simbol +
	fmt.Printf("angka negatif: %d\nangka positif: %d\n", negatifNum, positifNum)
}
