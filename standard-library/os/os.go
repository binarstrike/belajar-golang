package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args nilai pada argument saat binary hasil compile dijalankan
	args := os.Args

	// contoh: os.exe binar nugroho
	// atau
	// go run os.go binar nugroho
	// Args: []string{"/path/to/os.exe", "binar", "nugroho"}
	fmt.Printf("Args: %#v\n", args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Printf("Hostname: %s\n", hostname)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}
