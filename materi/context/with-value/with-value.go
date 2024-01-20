package main

import (
	"context"
	"fmt"

	"github.com/binarstrike/belajar-golang/pkg/person"
)

type MyKey int

const (
	Key1 MyKey = iota + 1
	Key2
	Key3
)

func main() {
	// ctx__1 := context.WithValue(context.Background(), Key1, &person.Person{FirstName: "Udin"})
	ctxParent := context.Background()
	ctx__1 := context.WithValue(ctxParent, Key1, &person.Person{FirstName: "Udin"})
	ctx__2 := context.WithValue(ctx__1, Key2, &person.Person{FirstName: "Dika"})
	ctx__3 := context.WithValue(ctxParent, Key1, &person.Person{FirstName: "Otong"})

	fmt.Printf("ctx__1 Key1: %#v\n", ctx__1.Value(Key1))
	fmt.Printf("ctx__2 Key1: %#v\n", ctx__2.Value(Key1))
	fmt.Printf("ctx__2 Key2: %#v\n", ctx__2.Value(Key2))

	fmt.Printf("ctx__3 Key1: %#v\n", ctx__3.Value(Key1))
	fmt.Printf("ctx__3 Key2: %#v\n", ctx__3.Value(Key2)) // nil

	if p1, ok := ctx__1.Value(Key1).(*person.Person); ok {
		p1.SetAge(17)
		p1.SetLastName("Alex")
		fmt.Printf("\nctx__1 Key1: %#v\n", ctx__1.Value(Key1)) // Udin Alex 17
	}
}
