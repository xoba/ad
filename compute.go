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
	a := pow(z, y) + pow(x, 2) + log(z+y*x*(exp(x)+x*y+x/y))
	fmt.Printf("formula: %f\n", a)
	c, grad := Compute(x, y, z)
	fmt.Printf("parsed : %f\n", c)
	fmt.Printf("diff   : %f\n", a-c)
	fmt.Printf("grad = %v\n", grad)
	delta := 0.000010
	tmp := a
	{
		z += delta
		a := pow(z, y) + pow(x, 2) + log(z+y*x*(exp(x)+x*y+x/y))
		z -= delta
		fmt.Printf("df/dz = %f\n", (a-tmp)/delta)
	}
	{
		y += delta
		a := pow(z, y) + pow(x, 2) + log(z+y*x*(exp(x)+x*y+x/y))
		y -= delta
		fmt.Printf("df/dy = %f\n", (a-tmp)/delta)
	}
	{
		x += delta
		a := pow(z, y) + pow(x, 2) + log(z+y*x*(exp(x)+x*y+x/y))
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
	switch i {
	case 0:
		return b * math.Pow(a, b-1)
	case 1:
		return math.Pow(a, b) * math.Log(a)
	default:
		panic("illegal index")
	}
}

func Compute(x, y, z float64) (float64, map[string]float64) {
	grad := make(map[string]float64)
	v0 := z
	v1 := y
	v2 := x
	s0 := pow(v0, v1)
	s1 := pow(v2, 2.000000)
	s2 := add(s0, s1)
	s3 := multiply(v1, v2)
	s4 := exp(v2)
	s5 := multiply(v2, v1)
	s6 := add(s4, s5)
	s7 := divide(v2, v1)
	s8 := add(s6, s7)
	s9 := multiply(s3, s8)
	s10 := add(v0, s9)
	s11 := log(s10)
	s12 := add(s2, s11)
	bs12 := 1.000000
	bs11 := 0.000000
	bs11 += bs12 * dadd(1, s2, s11)
	bs10 := 0.000000
	bs10 += bs11 * dlog(0, s10)
	bs9 := 0.000000
	bs9 += bs10 * dadd(1, v0, s9)
	bs8 := 0.000000
	bs8 += bs9 * dmultiply(1, s3, s8)
	bs7 := 0.000000
	bs7 += bs8 * dadd(1, s6, s7)
	bs6 := 0.000000
	bs6 += bs8 * dadd(0, s6, s7)
	bs5 := 0.000000
	bs5 += bs6 * dadd(1, s4, s5)
	bs4 := 0.000000
	bs4 += bs6 * dadd(0, s4, s5)
	bs3 := 0.000000
	bs3 += bs9 * dmultiply(0, s3, s8)
	bs2 := 0.000000
	bs2 += bs12 * dadd(0, s2, s11)
	bs1 := 0.000000
	bs1 += bs2 * dadd(1, s0, s1)
	bs0 := 0.000000
	bs0 += bs2 * dadd(0, s0, s1)
	bv2 := 0.000000
	bv2 += bs1 * dpow(0, v2, 2.000000)
	bv2 += bs3 * dmultiply(1, v1, v2)
	bv2 += bs4 * dexp(0, v2)
	bv2 += bs5 * dmultiply(0, v2, v1)
	bv2 += bs7 * ddivide(0, v2, v1)
	grad["x"] = bv2
	bv1 := 0.000000
	bv1 += bs0 * dpow(1, v0, v1)
	bv1 += bs3 * dmultiply(0, v1, v2)
	bv1 += bs5 * dmultiply(1, v2, v1)
	bv1 += bs7 * ddivide(1, v2, v1)
	grad["y"] = bv1
	bv0 := 0.000000
	bv0 += bs0 * dpow(0, v0, v1)
	bv0 += bs10 * dadd(0, v0, s9)
	grad["z"] = bv0
	return s12, grad
}
