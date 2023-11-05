package main

import "fmt"

// Label

// label digunakan untuk identifikasi sebuah statement seperti for..loop, switch..case dan pemanggilan sebuah fungsi
func main() {
outer_loop:
	for i := 0; true; i++ {
	inner_loop:
		for a := 0; true; a++ {
			if a == 2 {
				fmt.Println("break inner_loop")
				break inner_loop
			}
			fmt.Println("inner loop", i)
		}
		if i == 3 {
			fmt.Println("break outer_loop")
			break outer_loop
		}
		fmt.Println("outer loop", i)
	}

	// keyword goto digunakan untuk melompat atau mengalihakan jalannya eksekusi kode program, seperti contoh baris
	// kode dibawah dimana ketika sudah menjalankan fungsi fmt.Println() dengan parameter sebuah string "A" dan "C"
	// langsung dialihkan ke label `print_d` dimana akan menjalankan fungsi fmt.Println("D") yang mengeluarkan output pada
	// terminal berupa string "D" dan fungsi fmt.Println("C") akan diabaikan dan tidak pernah dijalankan.
	fmt.Println("A")
	fmt.Println("B")
	goto print_d
	fmt.Println("C")
print_d:
	fmt.Println("D")
}
