package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("running compute.go")
	rand.Seed(time.Now().UTC().UnixNano())
	x := rand.Float64()
	fmt.Printf("x = %f\n", x)
	y := rand.Float64()
	fmt.Printf("y = %f\n", y)
	z := rand.Float64()
	fmt.Printf("z = %f\n", z)
	a := log(z + y*x*(exp(x)+x*y+x/y))
	fmt.Printf("formula: %f\n", a)
	c, grad := Compute(x, y, z)
	fmt.Printf("parsed : %f\n", c)
	fmt.Printf("diff   : %f\n", a-c)
	fmt.Printf("grad = %v\n", grad)
	delta := 0.000010
	tmp := a
	{
		z += delta
		a := log(z + y*x*(exp(x)+x*y+x/y))
		z -= delta
		fmt.Printf("df/dz = %f\n", (a-tmp)/delta)
	}
	{
		y += delta
		a := log(z + y*x*(exp(x)+x*y+x/y))
		y -= delta
		fmt.Printf("df/dy = %f\n", (a-tmp)/delta)
	}
	{
		x += delta
		a := log(z + y*x*(exp(x)+x*y+x/y))
		x -= delta
		fmt.Printf("df/dx = %f\n", (a-tmp)/delta)
	}

}

func add(a, b float64) float64 {
	return a + b
}
func dadd(i int, a, b float64) float64 {
	return 1
}

func multiply(a, b float64) float64 {
	return a * b
}
func dmultiply(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b
	case 1:
		return a
	default:
		panic("illegal index")
	}
}

func subtract(a, b float64) float64 {
	return a - b
}
func dsubtract(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1
	case 1:
		return -1
	default:
		panic("illegal index")
	}
}

func divide(a, b float64) float64 {
	return a / b
}
func ddivide(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1 / b
	case 1:
		return -a / (b * b)
	default:
		panic("illegal index")
	}
	panic("unimplemented")
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}
func dsqrt(_int, a float64) float64 {
	return 0.5 * math.Pow(a, -1.5)
}

func exp(a float64) float64 {
	return math.Exp(a)
}
func dexp(_ int, a float64) float64 {
	return exp(a)
}

func log(a float64) float64 {
	return math.Log(a)
}
func dlog(_ int, a float64) float64 {
	return 1 / a
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
func dpow(i int, a, b float64) float64 {
	panic("unimplemented")
}

func Compute(x, y, z float64) (float64, map[string]float64) {
	grad := make(map[string]float64)
	v0 := z
	fmt.Printf("v0 = %f\n", v0)
	v1 := y
	fmt.Printf("v1 = %f\n", v1)
	v2 := x
	fmt.Printf("v2 = %f\n", v2)
	s0 := multiply(v1, v2)
	fmt.Printf("s0 = %f\n", s0)
	s1 := exp(v2)
	fmt.Printf("s1 = %f\n", s1)
	s2 := multiply(v2, v1)
	fmt.Printf("s2 = %f\n", s2)
	s3 := add(s1, s2)
	fmt.Printf("s3 = %f\n", s3)
	s4 := divide(v2, v1)
	fmt.Printf("s4 = %f\n", s4)
	s5 := add(s3, s4)
	fmt.Printf("s5 = %f\n", s5)
	s6 := multiply(s0, s5)
	fmt.Printf("s6 = %f\n", s6)
	s7 := add(v0, s6)
	fmt.Printf("s7 = %f\n", s7)
	s8 := log(s7)
	fmt.Printf("s8 = %f\n", s8)
	bs8 := 1.000000
	bs7 := 0.000000
	bs7 += bs8 * dlog(0, s7)
	bs6 := 0.000000
	bs6 += bs7 * dadd(1, v0, s6)
	bs5 := 0.000000
	bs5 += bs6 * dmultiply(1, s0, s5)
	bs4 := 0.000000
	bs4 += bs5 * dadd(1, s3, s4)
	bs3 := 0.000000
	bs3 += bs5 * dadd(0, s3, s4)
	bs2 := 0.000000
	bs2 += bs3 * dadd(1, s1, s2)
	bs1 := 0.000000
	bs1 += bs3 * dadd(0, s1, s2)
	bs0 := 0.000000
	bs0 += bs6 * dmultiply(0, s0, s5)
	bv2 := 0.000000
	bv2 += bs0 * dmultiply(1, v1, v2)
	bv2 += bs1 * dexp(0, v2)
	bv2 += bs2 * dmultiply(0, v2, v1)
	bv2 += bs4 * ddivide(0, v2, v1)
	grad["x"] = bv2
	bv1 := 0.000000
	bv1 += bs0 * dmultiply(0, v1, v2)
	bv1 += bs2 * dmultiply(1, v2, v1)
	bv1 += bs4 * ddivide(1, v2, v1)
	grad["y"] = bv1
	bv0 := 0.000000
	bv0 += bs7 * dadd(0, v0, s6)
	grad["z"] = bv0
	return s8, grad
}
