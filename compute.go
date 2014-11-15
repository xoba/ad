package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	x := rand.Float64()
	y := rand.Float64()
	z := rand.Float64()
	a := x + sqrt(pow(z, 3)) + 99*(-5*x+55)/6 + y
	fmt.Printf("formula: %f\n", a)
	fmt.Printf("parsed : %f\n", Compute(x, y, z))
	fmt.Printf("diff   : %f\n", a-Compute(x, y, z))
}

func Compute(x, y, z float64) float64 {
	v0 := x
	v1 := z
	v2 := y
	s0 := pow(v1, 3.000000)
	s1 := sqrt(s0)
	s2 := add(v0, s1)
	s3 := multiply(-1.000000, 5.000000)
	s4 := multiply(s3, v0)
	s5 := add(s4, 55.000000)
	s6 := multiply(99.000000, s5)
	s7 := divide(s6, 6.000000)
	s8 := add(s2, s7)
	s9 := add(s8, v2)
	return s9
}

func add(a, b float64) float64 {
	return a + b
}

func multiply(a, b float64) float64 {
	return a * b
}

func subtract(a, b float64) float64 {
	return a - b
}

func divide(a, b float64) float64 {
	return a / b
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
