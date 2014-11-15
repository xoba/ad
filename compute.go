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
	x0 := rand.Float64()
	fmt.Printf("x0 = %f\n", x0)
	x1 := rand.Float64()
	fmt.Printf("x1 = %f\n", x1)
	x2 := rand.Float64()
	fmt.Printf("x2 = %f\n", x2)
	x3 := rand.Float64()
	fmt.Printf("x3 = %f\n", x3)
	x4 := rand.Float64()
	fmt.Printf("x4 = %f\n", x4)
	x5 := rand.Float64()
	fmt.Printf("x5 = %f\n", x5)
	x6 := rand.Float64()
	fmt.Printf("x6 = %f\n", x6)
	x7 := rand.Float64()
	fmt.Printf("x7 = %f\n", x7)
	x8 := rand.Float64()
	fmt.Printf("x8 = %f\n", x8)
	x9 := rand.Float64()
	fmt.Printf("x9 = %f\n", x9)
	y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
	fmt.Printf("formula: %f\n", y)
	c, grad := Compute(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9)
	fmt.Printf("parsed : %f\n", c)
	fmt.Printf("diff   : %f\n", y-c)
	fmt.Printf("grad = %v\n", grad)
	delta := 0.000010
	tmp := y
	{
		x0 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x0 -= delta
		fmt.Printf("df/dx0 = %f\n", (y-tmp)/delta)
	}
	{
		x4 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x4 -= delta
		fmt.Printf("df/dx4 = %f\n", (y-tmp)/delta)
	}
	{
		x6 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x6 -= delta
		fmt.Printf("df/dx6 = %f\n", (y-tmp)/delta)
	}
	{
		x9 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x9 -= delta
		fmt.Printf("df/dx9 = %f\n", (y-tmp)/delta)
	}
	{
		x1 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x1 -= delta
		fmt.Printf("df/dx1 = %f\n", (y-tmp)/delta)
	}
	{
		x2 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x2 -= delta
		fmt.Printf("df/dx2 = %f\n", (y-tmp)/delta)
	}
	{
		x3 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x3 -= delta
		fmt.Printf("df/dx3 = %f\n", (y-tmp)/delta)
	}
	{
		x5 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x5 -= delta
		fmt.Printf("df/dx5 = %f\n", (y-tmp)/delta)
	}
	{
		x7 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x7 -= delta
		fmt.Printf("df/dx7 = %f\n", (y-tmp)/delta)
	}
	{
		x8 += delta
		y := 0.604660*x0 + 0.940509*x1 + 0.664560*x2 + 0.437714*x3 + 0.424637*x4 + 0.686823*x5 + 0.065637*x6 + 0.156519*x7 + 0.096970*x8 + 0.300912*x9
		x8 -= delta
		fmt.Printf("df/dx8 = %f\n", (y-tmp)/delta)
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

func Compute(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9 float64) (float64, map[string]float64) {
	grad := make(map[string]float64)
	v0 := x0
	v1 := x1
	v2 := x2
	v3 := x3
	v4 := x4
	v5 := x5
	v6 := x6
	v7 := x7
	v8 := x8
	v9 := x9
	s0 := multiply(0.604660, v0)
	s1 := multiply(0.940509, v1)
	s2 := add(s0, s1)
	s3 := multiply(0.664560, v2)
	s4 := add(s2, s3)
	s5 := multiply(0.437714, v3)
	s6 := add(s4, s5)
	s7 := multiply(0.424637, v4)
	s8 := add(s6, s7)
	s9 := multiply(0.686823, v5)
	s10 := add(s8, s9)
	s11 := multiply(0.065637, v6)
	s12 := add(s10, s11)
	s13 := multiply(0.156519, v7)
	s14 := add(s12, s13)
	s15 := multiply(0.096970, v8)
	s16 := add(s14, s15)
	s17 := multiply(0.300912, v9)
	s18 := add(s16, s17)
	bs18 := 1.000000
	bs17 := 0.000000
	bs17 += bs18 * dadd(1, s16, s17)
	bs16 := 0.000000
	bs16 += bs18 * dadd(0, s16, s17)
	bs15 := 0.000000
	bs15 += bs16 * dadd(1, s14, s15)
	bs14 := 0.000000
	bs14 += bs16 * dadd(0, s14, s15)
	bs13 := 0.000000
	bs13 += bs14 * dadd(1, s12, s13)
	bs12 := 0.000000
	bs12 += bs14 * dadd(0, s12, s13)
	bs11 := 0.000000
	bs11 += bs12 * dadd(1, s10, s11)
	bs10 := 0.000000
	bs10 += bs12 * dadd(0, s10, s11)
	bs9 := 0.000000
	bs9 += bs10 * dadd(1, s8, s9)
	bs8 := 0.000000
	bs8 += bs10 * dadd(0, s8, s9)
	bs7 := 0.000000
	bs7 += bs8 * dadd(1, s6, s7)
	bs6 := 0.000000
	bs6 += bs8 * dadd(0, s6, s7)
	bs5 := 0.000000
	bs5 += bs6 * dadd(1, s4, s5)
	bs4 := 0.000000
	bs4 += bs6 * dadd(0, s4, s5)
	bs3 := 0.000000
	bs3 += bs4 * dadd(1, s2, s3)
	bs2 := 0.000000
	bs2 += bs4 * dadd(0, s2, s3)
	bs1 := 0.000000
	bs1 += bs2 * dadd(1, s0, s1)
	bs0 := 0.000000
	bs0 += bs2 * dadd(0, s0, s1)
	bv9 := 0.000000
	bv9 += bs17 * dmultiply(1, 0.300912, v9)
	grad["x9"] = bv9
	bv8 := 0.000000
	bv8 += bs15 * dmultiply(1, 0.096970, v8)
	grad["x8"] = bv8
	bv7 := 0.000000
	bv7 += bs13 * dmultiply(1, 0.156519, v7)
	grad["x7"] = bv7
	bv6 := 0.000000
	bv6 += bs11 * dmultiply(1, 0.065637, v6)
	grad["x6"] = bv6
	bv5 := 0.000000
	bv5 += bs9 * dmultiply(1, 0.686823, v5)
	grad["x5"] = bv5
	bv4 := 0.000000
	bv4 += bs7 * dmultiply(1, 0.424637, v4)
	grad["x4"] = bv4
	bv3 := 0.000000
	bv3 += bs5 * dmultiply(1, 0.437714, v3)
	grad["x3"] = bv3
	bv2 := 0.000000
	bv2 += bs3 * dmultiply(1, 0.664560, v2)
	grad["x2"] = bv2
	bv1 := 0.000000
	bv1 += bs1 * dmultiply(1, 0.940509, v1)
	grad["x1"] = bv1
	bv0 := 0.000000
	bv0 += bs0 * dmultiply(1, 0.604660, v0)
	grad["x0"] = bv0
	return s18, grad
}
