package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/binarstrike/belajar-golang/pkg/person"
)

func main() {
	pChan := make(chan *person.Person, 20)

	p1 := &person.Person{FirstName: "Ucup", LastName: "James", Age: 10}

	wg := &sync.WaitGroup{}

	count := 10
	wg.Add(count)

	go func() {
		defer close(pChan)
		for i := 1; i <= count; i++ {
			// time.Sleep(time.Millisecond * 100)
			p1.Age++
			pChan <- p1
		}
	}()

	time.Sleep(1 * time.Second)

	go func() {
		// for a := range pChan {
		// 	fmt.Println(a.FirstName)
		// 	fmt.Println(a.LastName)
		// 	fmt.Println(a.Age)
		// 	wg.Done()
		// }
	}()

	for i := 1; i <= count; i++ {
		fmt.Println(<-pChan)
		wg.Done()
	}
	// wg.Wait()
	// fmt.Println(<-ax)
}
