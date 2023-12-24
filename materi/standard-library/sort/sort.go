package main

import (
	"fmt"
	"sort"

	p "github.com/binarstrike/belajar-golang/pkg/person"
)

type People []p.Person

// membuat method-method sesuai kontrak atau interface sort.Interface
// dan bila dilakukan pengurutan atau sorting maka akan diurutkan berdasarkan People[].Age
func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	var (
		p1 People = People{
			{FirstName: "Ucup", Age: 17},
			{FirstName: "Otong", Age: 21},
			{FirstName: "Udin", Age: 32},
			{FirstName: "Dika", Age: 18},
		}
		p2 []p.Person = p1
	)

	// sort.Slice() digunakan untuk mengurutkan data tanpa method-method pada sort.Interface
	sort.Slice(p1, func(i, j int) bool {
		return p1[i].Age < p1[j].Age
	})
	fmt.Println(p1)

	// sort.Sort() sama seperti sort.Slice() namun harus membuat method-method sesuai kontrak atau interface
	// sort.Interface pada slice
	// sort.Sort(p2) // ini akan error karena method-method berdasarkan dari kontrak interface sort.Interface
	// diimplementasikan oleh type alias People
	sort.Sort(People(p2))
	fmt.Println(p2)
}
