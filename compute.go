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
	s2 := rand.Float64()
	fmt.Printf("s2 = %f\n", s2)
	x0 := rand.Float64()
	fmt.Printf("x0 = %f\n", x0)
	x1 := rand.Float64()
	fmt.Printf("x1 = %f\n", x1)
	x2 := rand.Float64()
	fmt.Printf("x2 = %f\n", x2)
	y := rand.Float64()
	fmt.Printf("y = %f\n", y)

	c1, grad1 := ComputeAD(s2, x0, x1, x2, y)
	fmt.Printf("ad value: %f\n", c1)
	fmt.Printf("ad grad : %v\n", grad1)

	c2, grad2 := ComputeNum(s2, x0, x1, x2, y)
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

func xa8b32499_add(a, b float64) float64 {
	return a + b
}
func d_xa8b32499_add(i int, a, b float64) float64 {
	return 1
}

func xa8b32499_multiply(a, b float64) float64 {
	return a * b
}
func d_xa8b32499_multiply(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b
	case 1:
		return a
	default:
		panic("illegal index")
	}
}

func xa8b32499_subtract(a, b float64) float64 {
	return a - b
}
func d_xa8b32499_subtract(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1
	case 1:
		return -1
	default:
		panic("illegal index")
	}
}

func xa8b32499_divide(a, b float64) float64 {
	return a / b
}
func d_xa8b32499_divide(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1 / b
	case 1:
		return -a / (b * b)
	default:
		panic("illegal index")
	}
}

func xa8b32499_sqrt(a float64) float64 {
	return math.Sqrt(a)
}
func d_xa8b32499_sqrt(_int, a float64) float64 {
	return 0.5 * math.Pow(a, -1.5)
}

func xa8b32499_exp(a float64) float64 {
	return math.Exp(a)
}
func d_xa8b32499_exp(_ int, a float64) float64 {
	return math.Exp(a)
}

func xa8b32499_log(a float64) float64 {
	return math.Log(a)
}
func d_xa8b32499_log(_ int, a float64) float64 {
	return 1 / a
}

func xa8b32499_pow(a, b float64) float64 {
	return math.Pow(a, b)
}
func d_xa8b32499_pow(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b * math.Pow(a, b-1)
	case 1:
		return math.Pow(a, b) * math.Log(a)
	default:
		panic("illegal index")
	}
}

func log(a float64) float64 {
	return math.Log(a)
}
func exp(a float64) float64 {
	return math.Exp(a)
}

func ComputeAD(s2, x0, x1, x2, y float64) (float64, map[string]float64) {
	grad_xa8b32499 := make(map[string]float64)
	v_0_xa8b32499 := s2
	v_1_xa8b32499 := y
	v_2_xa8b32499 := x0
	v_3_xa8b32499 := x1
	v_4_xa8b32499 := x2
	s_0_xa8b32499 := xa8b32499_multiply(-1.000000, v_1_xa8b32499)
	s_1_xa8b32499 := xa8b32499_multiply(0.604660, v_2_xa8b32499)
	s_2_xa8b32499 := xa8b32499_multiply(0.940509, v_3_xa8b32499)
	s_3_xa8b32499 := xa8b32499_add(s_1_xa8b32499, s_2_xa8b32499)
	s_4_xa8b32499 := xa8b32499_multiply(0.664560, v_4_xa8b32499)
	s_5_xa8b32499 := xa8b32499_add(s_3_xa8b32499, s_4_xa8b32499)
	s_6_xa8b32499 := xa8b32499_multiply(s_0_xa8b32499, s_5_xa8b32499)
	s_7_xa8b32499 := xa8b32499_exp(s_6_xa8b32499)
	s_8_xa8b32499 := xa8b32499_add(1.000000, s_7_xa8b32499)
	s_9_xa8b32499 := xa8b32499_log(s_8_xa8b32499)
	s_10_xa8b32499 := xa8b32499_add(v_0_xa8b32499, s_9_xa8b32499)
	bs_10_xa8b32499 := 1.000000
	bs_9_xa8b32499 := 0.000000
	bs_9_xa8b32499 += bs_10_xa8b32499 * d_xa8b32499_add(1, v_0_xa8b32499, s_9_xa8b32499)
	bs_8_xa8b32499 := 0.000000
	bs_8_xa8b32499 += bs_9_xa8b32499 * d_xa8b32499_log(0, s_8_xa8b32499)
	bs_7_xa8b32499 := 0.000000
	bs_7_xa8b32499 += bs_8_xa8b32499 * d_xa8b32499_add(1, 1.000000, s_7_xa8b32499)
	bs_6_xa8b32499 := 0.000000
	bs_6_xa8b32499 += bs_7_xa8b32499 * d_xa8b32499_exp(0, s_6_xa8b32499)
	bs_5_xa8b32499 := 0.000000
	bs_5_xa8b32499 += bs_6_xa8b32499 * d_xa8b32499_multiply(1, s_0_xa8b32499, s_5_xa8b32499)
	bs_4_xa8b32499 := 0.000000
	bs_4_xa8b32499 += bs_5_xa8b32499 * d_xa8b32499_add(1, s_3_xa8b32499, s_4_xa8b32499)
	bs_3_xa8b32499 := 0.000000
	bs_3_xa8b32499 += bs_5_xa8b32499 * d_xa8b32499_add(0, s_3_xa8b32499, s_4_xa8b32499)
	bs_2_xa8b32499 := 0.000000
	bs_2_xa8b32499 += bs_3_xa8b32499 * d_xa8b32499_add(1, s_1_xa8b32499, s_2_xa8b32499)
	bs_1_xa8b32499 := 0.000000
	bs_1_xa8b32499 += bs_3_xa8b32499 * d_xa8b32499_add(0, s_1_xa8b32499, s_2_xa8b32499)
	bs_0_xa8b32499 := 0.000000
	bs_0_xa8b32499 += bs_6_xa8b32499 * d_xa8b32499_multiply(0, s_0_xa8b32499, s_5_xa8b32499)
	bv_4_xa8b32499 := 0.000000
	bv_4_xa8b32499 += bs_4_xa8b32499 * d_xa8b32499_multiply(1, 0.664560, v_4_xa8b32499)
	grad_xa8b32499["x2"] = bv_4_xa8b32499
	bv_3_xa8b32499 := 0.000000
	bv_3_xa8b32499 += bs_2_xa8b32499 * d_xa8b32499_multiply(1, 0.940509, v_3_xa8b32499)
	grad_xa8b32499["x1"] = bv_3_xa8b32499
	bv_2_xa8b32499 := 0.000000
	bv_2_xa8b32499 += bs_1_xa8b32499 * d_xa8b32499_multiply(1, 0.604660, v_2_xa8b32499)
	grad_xa8b32499["x0"] = bv_2_xa8b32499
	bv_1_xa8b32499 := 0.000000
	bv_1_xa8b32499 += bs_0_xa8b32499 * d_xa8b32499_multiply(1, -1.000000, v_1_xa8b32499)
	grad_xa8b32499["y"] = bv_1_xa8b32499
	bv_0_xa8b32499 := 0.000000
	bv_0_xa8b32499 += bs_10_xa8b32499 * d_xa8b32499_add(0, v_0_xa8b32499, s_9_xa8b32499)
	grad_xa8b32499["s2"] = bv_0_xa8b32499
	return s_10_xa8b32499, grad_xa8b32499
}

func ComputeNum(s2, x0, x1, x2, y float64) (float64, map[string]float64) {
	grad_xa8b32499 := make(map[string]float64)
	delta := 0.000010
	calc := func() float64 {
		f := s2 + log(1+exp(-y*(0.604660*x0+0.940509*x1+0.664560*x2)))
		return f
	}
	tmp1 := calc()
	{
		x0 += delta
		tmp2 := calc()
		x0 -= delta
		grad_xa8b32499["x0"] = (tmp2 - tmp1) / delta
	}
	{
		x1 += delta
		tmp2 := calc()
		x1 -= delta
		grad_xa8b32499["x1"] = (tmp2 - tmp1) / delta
	}
	{
		x2 += delta
		tmp2 := calc()
		x2 -= delta
		grad_xa8b32499["x2"] = (tmp2 - tmp1) / delta
	}
	{
		s2 += delta
		tmp2 := calc()
		s2 -= delta
		grad_xa8b32499["s2"] = (tmp2 - tmp1) / delta
	}
	{
		y += delta
		tmp2 := calc()
		y -= delta
		grad_xa8b32499["y"] = (tmp2 - tmp1) / delta
	}
	return tmp1, grad_xa8b32499
}
