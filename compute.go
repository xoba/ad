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
	v_0 := x
	v_1 := z
	v_2 := y
	v1 := pow(v_1, 3.000000)
	v2 := sqrt(v1)
	v3 := v_0 + v2
	v4 := 5.000000 * v_0
	v5 := v4 + 55.000000
	v6 := 99.000000 * v5
	v7 := v6 / 6.000000
	v8 := v3 + v7
	v9 := v8 + v_2
	return v9
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
