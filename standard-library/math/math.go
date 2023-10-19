package main

import (
	"math"

	a "github.com/binarstrike/belajar-golang/pkg/alias"
)

func main() {
	const floatNum1 float64 = 12.988897
	const floatNum2 float64 = 127.19248
	const floatNegativeNum float64 = -73.23213

	a.Print(a.F("math.Round: %g", math.Round(floatNum1)))
	a.Print(a.F("math.Ceil: %g", math.Ceil(floatNum2)))
	a.Print(a.F("math.Floor: %g", math.Floor(floatNum2)))
	a.Print(a.F("math.Max: %g", math.Max(math.Round(floatNum1), math.Round(floatNum2))))
	a.Print(a.F("math.Min: %g", math.Min(floatNum2, floatNum1)))
	a.Print(a.F("math.Log: %g", math.Log((math.Round(floatNum2)))))
	a.Print(a.F("math.Abs: %g", math.Abs(floatNegativeNum)))
}
