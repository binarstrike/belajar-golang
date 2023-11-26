package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Volume interface {
	constraints.Float | constraints.Integer
}

type Rider[T Volume] interface {
	MotorAttributes() (brand, model string, wheels uint, fuelCapacity, oilCapacity T)
}

type Motorcycle[T Volume] struct {
	brand                     string
	model                     string
	wheels                    uint
	fuelCapacity, oilCapacity T
}

func describeMotor[T Volume](r Rider[T]) {
	br, m, w, fc, oc := r.MotorAttributes()
	fmt.Printf(`
Brand: %s
Model: %s
Wheels: %d
Fuel Capacity: %v liter
Oil Capacity: %v liter
`, br, m, w, fc, oc)
}

func (m Motorcycle[T]) MotorAttributes() (brand, model string, wheels uint, fuelCapacity, oilCapacity T) {
	return m.brand, m.model, m.wheels, m.fuelCapacity, m.oilCapacity
}

func main() {
	motor1 := Motorcycle[float64]{
		brand:        "Yamaha",
		model:        "KZX-235",
		wheels:       2,
		fuelCapacity: 7.5,
		oilCapacity:  1.5,
	}

	describeMotor(motor1)
}
