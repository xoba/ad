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
	a := x + sqrt(pow(z, 3)) + 99*(5*x+55)/6 + y
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
	s2 := v0 + s1
	s3 := 5.000000 * v0
	s4 := s3 + 55.000000
	s5 := 99.000000 * s4
	s6 := s5 / 6.000000
	s7 := s2 + s6
	s8 := s7 + v2
	return s8
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
