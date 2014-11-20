// created 2014-11-20T22:50:28.468Z
// see https://github.com/xoba/ad

package nn

import (
	"fmt"
	"math"
)

// automatically compute the value and gradient of "f:=x0*x1*x2*x3*x4*x5*x6*x7*x8*x9"
func ComputeAD(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9 float64) (float64, map[string]float64) {
	grad_pvt := make(map[string]float64)
	v_0_pvt := x0
	v_1_pvt := x1
	v_2_pvt := x2
	v_3_pvt := x3
	v_4_pvt := x4
	v_5_pvt := x5
	v_6_pvt := x6
	v_7_pvt := x7
	v_8_pvt := x8
	v_9_pvt := x9
	s_0_pvt := multiply_pvt(v_0_pvt, v_1_pvt)
	s_1_pvt := multiply_pvt(s_0_pvt, v_2_pvt)
	s_2_pvt := multiply_pvt(s_1_pvt, v_3_pvt)
	s_3_pvt := multiply_pvt(s_2_pvt, v_4_pvt)
	s_4_pvt := multiply_pvt(s_3_pvt, v_5_pvt)
	s_5_pvt := multiply_pvt(s_4_pvt, v_6_pvt)
	s_6_pvt := multiply_pvt(s_5_pvt, v_7_pvt)
	s_7_pvt := multiply_pvt(s_6_pvt, v_8_pvt)
	s_8_pvt := multiply_pvt(s_7_pvt, v_9_pvt)
	const b_s_8_pvt = 1.0
	b_s_7_pvt := 0.0
	b_s_7_pvt += b_s_8_pvt * (d_multiply_pvt(0, s_7_pvt, v_9_pvt))
	b_s_6_pvt := 0.0
	b_s_6_pvt += b_s_7_pvt * (d_multiply_pvt(0, s_6_pvt, v_8_pvt))
	b_s_5_pvt := 0.0
	b_s_5_pvt += b_s_6_pvt * (d_multiply_pvt(0, s_5_pvt, v_7_pvt))
	b_s_4_pvt := 0.0
	b_s_4_pvt += b_s_5_pvt * (d_multiply_pvt(0, s_4_pvt, v_6_pvt))
	b_s_3_pvt := 0.0
	b_s_3_pvt += b_s_4_pvt * (d_multiply_pvt(0, s_3_pvt, v_5_pvt))
	b_s_2_pvt := 0.0
	b_s_2_pvt += b_s_3_pvt * (d_multiply_pvt(0, s_2_pvt, v_4_pvt))
	b_s_1_pvt := 0.0
	b_s_1_pvt += b_s_2_pvt * (d_multiply_pvt(0, s_1_pvt, v_3_pvt))
	b_s_0_pvt := 0.0
	b_s_0_pvt += b_s_1_pvt * (d_multiply_pvt(0, s_0_pvt, v_2_pvt))
	b_v_9_pvt := 0.0
	b_v_9_pvt += b_s_8_pvt * (d_multiply_pvt(1, s_7_pvt, v_9_pvt))
	grad_pvt["x9"] = b_v_9_pvt
	b_v_8_pvt := 0.0
	b_v_8_pvt += b_s_7_pvt * (d_multiply_pvt(1, s_6_pvt, v_8_pvt))
	grad_pvt["x8"] = b_v_8_pvt
	b_v_7_pvt := 0.0
	b_v_7_pvt += b_s_6_pvt * (d_multiply_pvt(1, s_5_pvt, v_7_pvt))
	grad_pvt["x7"] = b_v_7_pvt
	b_v_6_pvt := 0.0
	b_v_6_pvt += b_s_5_pvt * (d_multiply_pvt(1, s_4_pvt, v_6_pvt))
	grad_pvt["x6"] = b_v_6_pvt
	b_v_5_pvt := 0.0
	b_v_5_pvt += b_s_4_pvt * (d_multiply_pvt(1, s_3_pvt, v_5_pvt))
	grad_pvt["x5"] = b_v_5_pvt
	b_v_4_pvt := 0.0
	b_v_4_pvt += b_s_3_pvt * (d_multiply_pvt(1, s_2_pvt, v_4_pvt))
	grad_pvt["x4"] = b_v_4_pvt
	b_v_3_pvt := 0.0
	b_v_3_pvt += b_s_2_pvt * (d_multiply_pvt(1, s_1_pvt, v_3_pvt))
	grad_pvt["x3"] = b_v_3_pvt
	b_v_2_pvt := 0.0
	b_v_2_pvt += b_s_1_pvt * (d_multiply_pvt(1, s_0_pvt, v_2_pvt))
	grad_pvt["x2"] = b_v_2_pvt
	b_v_1_pvt := 0.0
	b_v_1_pvt += b_s_0_pvt * (d_multiply_pvt(1, v_0_pvt, v_1_pvt))
	grad_pvt["x1"] = b_v_1_pvt
	b_v_0_pvt := 0.0
	b_v_0_pvt += b_s_0_pvt * (d_multiply_pvt(0, v_0_pvt, v_1_pvt))
	grad_pvt["x0"] = b_v_0_pvt
	return s_8_pvt, grad_pvt
}

func main() {
	fmt.Printf("running autodiff code of 2014-11-20T22:50:28.468Z on %q\n\n", "f:=x0*x1*x2*x3*x4*x5*x6*x7*x8*x9")
	x0 := -1.74629981957221058764
	fmt.Printf("setting x0 = %+.20f\n", x0)
	x1 := -1.67116136509261092868
	fmt.Printf("setting x1 = %+.20f\n", x1)
	x2 := -0.51383334065677832569
	fmt.Printf("setting x2 = %+.20f\n", x2)
	x3 := 0.27872203908858056431
	fmt.Printf("setting x3 = %+.20f\n", x3)
	x4 := 0.05263745025135846412
	fmt.Printf("setting x4 = %+.20f\n", x4)
	x5 := 0.21823102250635140198
	fmt.Printf("setting x5 = %+.20f\n", x5)
	x6 := 0.19188568240941417109
	fmt.Printf("setting x6 = %+.20f\n", x6)
	x7 := -1.15570674051818023109
	fmt.Printf("setting x7 = %+.20f\n", x7)
	x8 := 1.90997687280961336853
	fmt.Printf("setting x8 = %+.20f\n", x8)
	x9 := -0.55326262285648875050
	fmt.Printf("setting x9 = %+.20f\n", x9)
	fmt.Println()

	c1, grad1 := ComputeAD(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9)
	fmt.Printf("autodiff value   : %.20f\n", c1)
	fmt.Printf("autodiff gradient: %v\n\n", grad1)

	c2, grad2 := ComputeNumerical(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9)
	fmt.Printf("numeric value   : %.20f\n", c2)
	fmt.Printf("numeric gradient: %v\n\n", grad2)

	var total float64
	add := func(n string, x float64) {
		fmt.Printf("%s difference: %+.20f\n", n, x)
		total += math.Abs(x)
	}
	add("value", c1-c2)
	for k, v := range grad2 {
		add(fmt.Sprintf("grad[%3s]", k), grad1[k]-v)
	}
	fmt.Printf("\nsum of absolute differences: %+.20f\n", total)
}

// numerically compute the value and gradient of "f:=x0*x1*x2*x3*x4*x5*x6*x7*x8*x9"
func ComputeNumerical(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9 float64) (float64, map[string]float64) {
	grad_pvt := make(map[string]float64)
	const delta_pvt = 0.000010
	calc_pvt := func() float64 {
		f := x0 * x1 * x2 * x3 * x4 * x5 * x6 * x7 * x8 * x9
		return f
	}
	tmp1_pvt := calc_pvt()
	{
		x5 += delta_pvt
		tmp2_pvt := calc_pvt()
		x5 -= delta_pvt
		grad_pvt["x5"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x7 += delta_pvt
		tmp2_pvt := calc_pvt()
		x7 -= delta_pvt
		grad_pvt["x7"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x8 += delta_pvt
		tmp2_pvt := calc_pvt()
		x8 -= delta_pvt
		grad_pvt["x8"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x1 += delta_pvt
		tmp2_pvt := calc_pvt()
		x1 -= delta_pvt
		grad_pvt["x1"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x2 += delta_pvt
		tmp2_pvt := calc_pvt()
		x2 -= delta_pvt
		grad_pvt["x2"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x4 += delta_pvt
		tmp2_pvt := calc_pvt()
		x4 -= delta_pvt
		grad_pvt["x4"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x6 += delta_pvt
		tmp2_pvt := calc_pvt()
		x6 -= delta_pvt
		grad_pvt["x6"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x9 += delta_pvt
		tmp2_pvt := calc_pvt()
		x9 -= delta_pvt
		grad_pvt["x9"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x0 += delta_pvt
		tmp2_pvt := calc_pvt()
		x0 -= delta_pvt
		grad_pvt["x0"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x3 += delta_pvt
		tmp2_pvt := calc_pvt()
		x3 -= delta_pvt
		grad_pvt["x3"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	return tmp1_pvt, grad_pvt
}

func tan(a float64) float64 {
	return math.Tan(a)
}

func tan_pvt(a float64) float64 {
	return math.Tan(a)
}

func d_tan(a float64) float64 {
	return math.Pow(1/math.Cos(a), 2)
}

func d_tan_pvt(a float64) float64 {
	return math.Pow(1/math.Cos(a), 2)
}

func abs(a float64) float64 {
	return math.Abs(a)
}

func abs_pvt(a float64) float64 {
	return math.Abs(a)
}

func d_abs(a float64) float64 {
	switch {
	case a > 0:
		return +1
	case a < 0:
		return -1
	default:
		panic("illegal derivative when abs(0)=0")
	}
}

func d_abs_pvt(a float64) float64 {
	switch {
	case a > 0:
		return +1
	case a < 0:
		return -1
	default:
		panic("illegal derivative when abs(0)=0")
	}
}

func atan(a float64) float64 {
	return math.Atan(a)
}

func atan_pvt(a float64) float64 {
	return math.Atan(a)
}

func d_atan(a float64) float64 {
	return 1 / (1 + a*a)
}

func d_atan_pvt(a float64) float64 {
	return 1 / (1 + a*a)
}

func tanh(a float64) float64 {
	return math.Tanh(a)
}

func tanh_pvt(a float64) float64 {
	return math.Tanh(a)
}

func d_tanh(a float64) float64 {
	return 1 - math.Pow(math.Tanh(a), 2)
}

func d_tanh_pvt(a float64) float64 {
	return 1 - math.Pow(math.Tanh(a), 2)
}

func sin(a float64) float64 {
	return math.Sin(a)
}

func sin_pvt(a float64) float64 {
	return math.Sin(a)
}

func d_sin(a float64) float64 {
	return math.Cos(a)
}

func d_sin_pvt(a float64) float64 {
	return math.Cos(a)
}

func asin(a float64) float64 {
	return math.Asin(a)
}

func asin_pvt(a float64) float64 {
	return math.Asin(a)
}

func d_asin(a float64) float64 {
	return 1 / math.Sqrt(1-a*a)
}

func d_asin_pvt(a float64) float64 {
	return 1 / math.Sqrt(1-a*a)
}

func sinh(a float64) float64 {
	return math.Sinh(a)
}

func sinh_pvt(a float64) float64 {
	return math.Sinh(a)
}

func d_sinh(a float64) float64 {
	return math.Cosh(a)
}

func d_sinh_pvt(a float64) float64 {
	return math.Cosh(a)
}

func cos(a float64) float64 {
	return math.Cos(a)
}

func cos_pvt(a float64) float64 {
	return math.Cos(a)
}

func d_cos(a float64) float64 {
	return -math.Sin(a)
}

func d_cos_pvt(a float64) float64 {
	return -math.Sin(a)
}

func acos(a float64) float64 {
	return math.Acos(a)
}

func acos_pvt(a float64) float64 {
	return math.Acos(a)
}

func d_acos(a float64) float64 {
	return -1 / math.Sqrt(1-a*a)
}

func d_acos_pvt(a float64) float64 {
	return -1 / math.Sqrt(1-a*a)
}

func cosh(a float64) float64 {
	return math.Cosh(a)
}

func cosh_pvt(a float64) float64 {
	return math.Cosh(a)
}

func d_cosh(a float64) float64 {
	return math.Sinh(a)
}

func d_cosh_pvt(a float64) float64 {
	return math.Sinh(a)
}

func add(a, b float64) float64 {
	return a + b
}

func add_pvt(a, b float64) float64 {
	return a + b
}

func d_add(i int, a, b float64) float64 {
	return 1
}

func d_add_pvt(i int, a, b float64) float64 {
	return 1
}

func multiply(a, b float64) float64 {
	return a * b
}

func multiply_pvt(a, b float64) float64 {
	return a * b
}

func d_multiply(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b
	case 1:
		return a
	default:
		panic("illegal index")
	}
}

func d_multiply_pvt(i int, a, b float64) float64 {
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

func subtract_pvt(a, b float64) float64 {
	return a - b
}

func d_subtract(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1
	case 1:
		return -1
	default:
		panic("illegal index")
	}
}

func d_subtract_pvt(i int, a, b float64) float64 {
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

func divide_pvt(a, b float64) float64 {
	return a / b
}

func d_divide(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1 / b
	case 1:
		return -a / (b * b)
	default:
		panic("illegal index")
	}
}

func d_divide_pvt(i int, a, b float64) float64 {
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

func sqrt_pvt(a float64) float64 {
	return math.Sqrt(a)
}

func d_sqrt(a float64) float64 {
	return 0.5 * math.Pow(a, -0.5)
}

func d_sqrt_pvt(a float64) float64 {
	return 0.5 * math.Pow(a, -0.5)
}

func exp(a float64) float64 {
	return math.Exp(a)
}

func exp_pvt(a float64) float64 {
	return math.Exp(a)
}

func d_exp(a float64) float64 {
	return math.Exp(a)
}

func d_exp_pvt(a float64) float64 {
	return math.Exp(a)
}

func log(a float64) float64 {
	return math.Log(a)
}

func log_pvt(a float64) float64 {
	return math.Log(a)
}

func d_log(a float64) float64 {
	return 1 / a
}

func d_log_pvt(a float64) float64 {
	return 1 / a
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func pow_pvt(a, b float64) float64 {
	return math.Pow(a, b)
}

func d_pow(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b * math.Pow(a, b-1)
	case 1:
		return math.Pow(a, b) * math.Log(a)
	default:
		panic("illegal index")
	}
}

func d_pow_pvt(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b * math.Pow(a, b-1)
	case 1:
		return math.Pow(a, b) * math.Log(a)
	default:
		panic("illegal index")
	}
}
