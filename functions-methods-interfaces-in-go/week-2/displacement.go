package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	DisplacemntOfT := func(t float64) float64 {
		return ((0.5)*a*math.Pow(t, 2) + v0*t + s0)
	}
	return DisplacemntOfT
}

func main() {

	var a, v0, s0, t float64

	fmt.Print("Enter value for acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter value for initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("Enter value for initial displacement: ")
	fmt.Scan(&s0)

	fmt.Print("Enter value for elapsed time: ")
	fmt.Scan(&t)

	f := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement after %.2f seconds = %.2f\n", t, f(t))
	
}
