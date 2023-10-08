package main

import "fmt"

func checkAdmin(name string, check func(p1 string) bool) bool {
	return check(name)
}

func main() {
	admins := [...]string{"Ucup", "Dika", "Ani"}

	// fungsi yang tidak memiliki nama dan disimpan di sebuah variabel
	isAdmin := func(name string) bool {
		for _, val := range admins {
			if val == name {
				return true
			}
		}
		return false
	}

	fmt.Println(checkAdmin("Dika", isAdmin))  // true
	fmt.Println(checkAdmin("Joko", isAdmin))  // false
	fmt.Println(checkAdmin("Otong", isAdmin)) // false
	fmt.Println(checkAdmin("Ani", isAdmin))   // true
}
