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
	s1 := multiply(0.604660, v1)
	s2 := multiply(0.940509, v2)
	s3 := add(s1, s2)
	s4 := multiply(0.664560, v3)
	s5 := add(s3, s4)
	s6 := multiply(0.437714, v4)
	s7 := add(s5, s6)
	s8 := multiply(0.424637, v5)
	s9 := add(s7, s8)
	s10 := multiply(0.686823, v6)
	s11 := add(s9, s10)
	s12 := multiply(0.065637, v7)
	s13 := add(s11, s12)
	s14 := multiply(0.156519, v8)
	s15 := add(s13, s14)
	s16 := multiply(0.096970, v9)
	s17 := add(s15, s16)
	s18 := multiply(0.300912, v10)
	s19 := add(s17, s18)
	s20 := multiply(s0, s19)
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
	bs19 += bs20 * dmultiply(1, s0, s19)
	bs18 := 0.000000
	bs18 += bs19 * dadd(1, s17, s18)
	bs17 := 0.000000
	bs17 += bs19 * dadd(0, s17, s18)
	bs16 := 0.000000
	bs16 += bs17 * dadd(1, s15, s16)
	bs15 := 0.000000
	bs15 += bs17 * dadd(0, s15, s16)
	bs14 := 0.000000
	bs14 += bs15 * dadd(1, s13, s14)
	bs13 := 0.000000
	bs13 += bs15 * dadd(0, s13, s14)
	bs12 := 0.000000
	bs12 += bs13 * dadd(1, s11, s12)
	bs11 := 0.000000
	bs11 += bs13 * dadd(0, s11, s12)
	bs10 := 0.000000
	bs10 += bs11 * dadd(1, s9, s10)
	bs9 := 0.000000
	bs9 += bs11 * dadd(0, s9, s10)
	bs8 := 0.000000
	bs8 += bs9 * dadd(1, s7, s8)
	bs7 := 0.000000
	bs7 += bs9 * dadd(0, s7, s8)
	bs6 := 0.000000
	bs6 += bs7 * dadd(1, s5, s6)
	bs5 := 0.000000
	bs5 += bs7 * dadd(0, s5, s6)
	bs4 := 0.000000
	bs4 += bs5 * dadd(1, s3, s4)
	bs3 := 0.000000
	bs3 += bs5 * dadd(0, s3, s4)
	bs2 := 0.000000
	bs2 += bs3 * dadd(1, s1, s2)
	bs1 := 0.000000
	bs1 += bs3 * dadd(0, s1, s2)
	bs0 := 0.000000
	bs0 += bs20 * dmultiply(0, s0, s19)
	bv10 := 0.000000
	bv10 += bs18 * dmultiply(1, 0.300912, v10)
	grad["x9"] = bv10
	bv9 := 0.000000
	bv9 += bs16 * dmultiply(1, 0.096970, v9)
	grad["x8"] = bv9
	bv8 := 0.000000
	bv8 += bs14 * dmultiply(1, 0.156519, v8)
	grad["x7"] = bv8
	bv7 := 0.000000
	bv7 += bs12 * dmultiply(1, 0.065637, v7)
	grad["x6"] = bv7
	bv6 := 0.000000
	bv6 += bs10 * dmultiply(1, 0.686823, v6)
	grad["x5"] = bv6
	bv5 := 0.000000
	bv5 += bs8 * dmultiply(1, 0.424637, v5)
	grad["x4"] = bv5
	bv4 := 0.000000
	bv4 += bs6 * dmultiply(1, 0.437714, v4)
	grad["x3"] = bv4
	bv3 := 0.000000
	bv3 += bs4 * dmultiply(1, 0.664560, v3)
	grad["x2"] = bv3
	bv2 := 0.000000
	bv2 += bs2 * dmultiply(1, 0.940509, v2)
	grad["x1"] = bv2
	bv1 := 0.000000
	bv1 += bs1 * dmultiply(1, 0.604660, v1)
	grad["x0"] = bv1
	bv0 := 0.000000
	bv0 += bs0 * dmultiply(1, -1.000000, v0)
	grad["y"] = bv0
	return s23, grad
}

func ComputeNum(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, y float64) (float64, map[string]float64) {
	grad := make(map[string]float64)
	delta := 0.000010
	calc := func() float64 {
		f := log(1 + exp(-y*(0.604660*x0+0.940509*x1+0.664560*x2+0.437714*x3+0.424637*x4+0.686823*x5+0.065637*x6+0.156519*x7+0.096970*x8+0.300912*x9)))
		return f
	}
	tmp1 := calc()
	{
		x7 += delta
		tmp2 := calc()
		x7 -= delta
		grad["x7"] = (tmp2 - tmp1) / delta
	}
	{
		x9 += delta
		tmp2 := calc()
		x9 -= delta
		grad["x9"] = (tmp2 - tmp1) / delta
	}
	{
		y += delta
		tmp2 := calc()
		y -= delta
		grad["y"] = (tmp2 - tmp1) / delta
	}
	{
		x0 += delta
		tmp2 := calc()
		x0 -= delta
		grad["x0"] = (tmp2 - tmp1) / delta
	}
	{
		x1 += delta
		tmp2 := calc()
		x1 -= delta
		grad["x1"] = (tmp2 - tmp1) / delta
	}
	{
		x2 += delta
		tmp2 := calc()
		x2 -= delta
		grad["x2"] = (tmp2 - tmp1) / delta
	}
	{
		x3 += delta
		tmp2 := calc()
		x3 -= delta
		grad["x3"] = (tmp2 - tmp1) / delta
	}
	{
		x4 += delta
		tmp2 := calc()
		x4 -= delta
		grad["x4"] = (tmp2 - tmp1) / delta
	}
	{
		x5 += delta
		tmp2 := calc()
		x5 -= delta
		grad["x5"] = (tmp2 - tmp1) / delta
	}
	{
		x6 += delta
		tmp2 := calc()
		x6 -= delta
		grad["x6"] = (tmp2 - tmp1) / delta
	}
	{
		x8 += delta
		tmp2 := calc()
		x8 -= delta
		grad["x8"] = (tmp2 - tmp1) / delta
	}
	return tmp1, grad
}
