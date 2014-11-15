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
	a := x * (x + y*y)
	fmt.Printf("formula: %f\n", a)
	c := Compute(x, y)
	fmt.Printf("parsed : %f\n", c)
	fmt.Printf("diff   : %f\n", a-c)

	delta := 0.000010
	tmp := a
	{
		x += delta
		a := x * (x + y*y)
		x -= delta
		fmt.Printf("df/dx = %f\n", (a-tmp)/delta)
	}
	{
		y += delta
		a := x * (x + y*y)
		y -= delta
		fmt.Printf("df/dy = %f\n", (a-tmp)/delta)
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
func dsqrt(a float64) float64 {
	return 0.5 * math.Pow(a, -1.5)
}

func exp(a float64) float64 {
	return math.Exp(a)
}
func dexp(a float64) float64 {
	return exp(a)
}

func log(a float64) float64 {
	return math.Log(a)
}
func dlog(a float64) float64 {
	return 1 / a
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
func dpow(i int, a, b float64) float64 {
	panic("unimplemented")
}

func Compute(x, y float64) float64 {
	v0 := x
	fmt.Printf("v0 = %f\n", v0)
	v1 := y
	fmt.Printf("v1 = %f\n", v1)
	s0 := multiply(v1, v1)
	fmt.Printf("s0 = %f\n", s0)
	s1 := add(v0, s0)
	fmt.Printf("s1 = %f\n", s1)
	s2 := multiply(v0, s1)
	fmt.Printf("s2 = %f\n", s2)
	bs2 := 1.000000
	fmt.Printf("bs2 = %f\n", bs2)
	bs1 := 0.000000
	// bs1 += bs2 * ds2 / ds1 (multiply(v0,s1))
	fmt.Printf("bs1 = %f\n", bs1)
	bs0 := 0.000000
	// bs0 += bs1 * ds1 / ds0 (add(v0,s0))
	// bs0 += bs2 * ds2 / ds0 (multiply(v0,s1))
	fmt.Printf("bs0 = %f\n", bs0)
	bv1 := 0.000000
	// bv1 += bs0 * ds0 / dv1 (multiply(v1,v1))
	// bv1 += bs1 * ds1 / dv1 (add(v0,s0))
	// bv1 += bs2 * ds2 / dv1 (multiply(v0,s1))
	fmt.Printf("bv1 = %f\n", bv1)
	bv0 := 0.000000
	// bv0 += bv1 * dv1 / dv0 (y)
	// bv0 += bs0 * ds0 / dv0 (multiply(v1,v1))
	// bv0 += bs1 * ds1 / dv0 (add(v0,s0))
	// bv0 += bs2 * ds2 / dv0 (multiply(v0,s1))
	fmt.Printf("bv0 = %f\n", bv0)
	return s2
}
