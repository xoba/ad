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
	y := rand.Float64()
	fmt.Printf("y = %f\n", y)

	c1, grad1 := ComputeAD(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, y)
	fmt.Printf("ad value: %f\n", c1)
	fmt.Printf("ad grad : %v\n", grad1)

	c2, grad2 := ComputeNum(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, y)
	fmt.Printf("num value: %f\n", c2)
	fmt.Printf("num grad : %v\n", grad2)

	var total float64
	add := func(n string, x float64) {
		fmt.Printf("diff %s: %f\n", n, x)
		total += math.Abs(x)
	}
	add("value", c1-c2)
	for k, v := range grad2 {
		add(fmt.Sprintf("grad[%3s]", k), grad1[k]-v)
	}
	fmt.Printf("*** total diffs: %f\n", total)
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

func ComputeAD(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, y float64) (float64, map[string]float64) {
	grad := make(map[string]float64)
	v0 := y
	v1 := x0
	v2 := x1
	v3 := x2
	v4 := x3
	v5 := x4
	v6 := x5
	v7 := x6
	v8 := x7
	v9 := x8
	v10 := x9
	s0 := multiply(-1.000000, v0)
	s1 := multiply(s0, 0.604660)
	s2 := multiply(s1, v1)
	s3 := multiply(0.940509, v2)
	s4 := add(s2, s3)
	s5 := multiply(0.664560, v3)
	s6 := add(s4, s5)
	s7 := multiply(0.437714, v4)
	s8 := add(s6, s7)
	s9 := multiply(0.424637, v5)
	s10 := add(s8, s9)
	s11 := multiply(0.686823, v6)
	s12 := add(s10, s11)
	s13 := multiply(0.065637, v7)
	s14 := add(s12, s13)
	s15 := multiply(0.156519, v8)
	s16 := add(s14, s15)
	s17 := multiply(0.096970, v9)
	s18 := add(s16, s17)
	s19 := multiply(0.300912, v10)
	s20 := add(s18, s19)
	s21 := exp(s20)
	s22 := add(1.000000, s21)
	s23 := log(s22)
	bs23 := 1.000000
	bs22 := 0.000000
	bs22 += bs23 * dlog(0, s22)
	bs21 := 0.000000
	bs21 += bs22 * dadd(1, 1.000000, s21)
	bs20 := 0.000000
	bs20 += bs21 * dexp(0, s20)
	bs19 := 0.000000
	bs19 += bs20 * dadd(1, s18, s19)
	bs18 := 0.000000
	bs18 += bs20 * dadd(0, s18, s19)
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
	bs1 += bs2 * dmultiply(0, s1, v1)
	bs0 := 0.000000
	bs0 += bs1 * dmultiply(0, s0, 0.604660)
	bv10 := 0.000000
	bv10 += bs19 * dmultiply(1, 0.300912, v10)
	grad["x9"] = bv10
	bv9 := 0.000000
	bv9 += bs17 * dmultiply(1, 0.096970, v9)
	grad["x8"] = bv9
	bv8 := 0.000000
	bv8 += bs15 * dmultiply(1, 0.156519, v8)
	grad["x7"] = bv8
	bv7 := 0.000000
	bv7 += bs13 * dmultiply(1, 0.065637, v7)
	grad["x6"] = bv7
	bv6 := 0.000000
	bv6 += bs11 * dmultiply(1, 0.686823, v6)
	grad["x5"] = bv6
	bv5 := 0.000000
	bv5 += bs9 * dmultiply(1, 0.424637, v5)
	grad["x4"] = bv5
	bv4 := 0.000000
	bv4 += bs7 * dmultiply(1, 0.437714, v4)
	grad["x3"] = bv4
	bv3 := 0.000000
	bv3 += bs5 * dmultiply(1, 0.664560, v3)
	grad["x2"] = bv3
	bv2 := 0.000000
	bv2 += bs3 * dmultiply(1, 0.940509, v2)
	grad["x1"] = bv2
	bv1 := 0.000000
	bv1 += bs2 * dmultiply(1, s1, v1)
	grad["x0"] = bv1
	bv0 := 0.000000
	bv0 += bs0 * dmultiply(1, -1.000000, v0)
	grad["y"] = bv0
	return s23, grad
}

func ComputeNum(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, y float64) (float64, map[string]float64) {
	grad := make(map[string]float64)
	delta := 0.000010
	f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
	tmp := f
	{
		x9 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x9 -= delta
		grad["x9"] = (f - tmp) / delta
	}
	{
		y += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		y -= delta
		grad["y"] = (f - tmp) / delta
	}
	{
		x2 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x2 -= delta
		grad["x2"] = (f - tmp) / delta
	}
	{
		x3 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x3 -= delta
		grad["x3"] = (f - tmp) / delta
	}
	{
		x5 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x5 -= delta
		grad["x5"] = (f - tmp) / delta
	}
	{
		x6 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x6 -= delta
		grad["x6"] = (f - tmp) / delta
	}
	{
		x7 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x7 -= delta
		grad["x7"] = (f - tmp) / delta
	}
	{
		x8 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x8 -= delta
		grad["x8"] = (f - tmp) / delta
	}
	{
		x0 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x0 -= delta
		grad["x0"] = (f - tmp) / delta
	}
	{
		x1 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x1 -= delta
		grad["x1"] = (f - tmp) / delta
	}
	{
		x4 += delta
		f := log(1 + exp(-y*0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9))
		x4 -= delta
		grad["x4"] = (f - tmp) / delta
	}
	return f, grad
}
