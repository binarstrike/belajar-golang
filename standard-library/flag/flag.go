package main

import (
	"flag"
	"fmt"
)

func main() {
	var isDahMakan bool

	// arg := flag.String("flag name", "default value", "flag description")
	host := flag.String("host", "localhost", "Hostname")
	user := flag.String("user", "", "Username")
	password := flag.String("password", "", "Password")
	flag.BoolVar(&isDahMakan, "isDahMakan", false, "Udah makan?")

	flag.Parse()

	fmt.Printf("Host: %s\nUsername: %s\nPassword: %s\n", *host, *user, *password)

	if isDahMakan {
		fmt.Println(*user, "sudah makan")
	} else {
		fmt.Println(*user, "belum makan")
	}
	/*
		go run .\standard-library\flag\flag.go -user budi --password=idub

		output:

		Host: localhost
		Username: budi
		Password: idub
		budi belum makan
	*/
}
