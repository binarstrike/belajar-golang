package main

import (
	"fmt"

	"github.com/binarstrike/belajar-golang/pkg/hero"
	"github.com/binarstrike/belajar-golang/pkg/person"
)

func main() {
	p1 := &person.Person{FirstName: "Ucup", LastName: "Santoso", Age: 18}

	fmt.Printf("p1: %#v\n", p1)
	p1.SetPersonFirstName("Albert")
	fmt.Printf("p1: %#v\n", p1)
	p1.SetPersonLastName("Supradi")
	fmt.Printf("p1: %#v\n", p1)
	p1.SetPersonAge(20)
	fmt.Printf("p1: %#v\n\n", p1)

	miya := &hero.Hero{
		Name:         "Miya",
		Role:         hero.RoleMarksman,
		AttackPoint:  100,
		HealthPoint:  1000,
		DefensePoint: 300,
	}
	fmt.Printf("hero Miya: %#v\n", miya)
	DisplayHeroAbility(miya)
}

func DisplayHeroAbility(hero hero.IHero) {
	attack, hp, defense := hero.GetHeroAbility()
	fmt.Printf("Attack: %d\nHealthPoint: %d\nDefense: %d\n", attack, hp, defense)
}
