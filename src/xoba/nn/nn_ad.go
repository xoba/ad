// autogenerated, do not edit!
// see https://github.com/xoba/ad

package nn

import (
	"math"
)

// automatically compute the value and gradient of "f := log( 1 + exp(-y * (b0 +  b1 * x1 + b2 * y1)))\n"
func ComputeAD(b0, b1, b2, x1, y, y1 float64) (float64, map[string]float64) {
	grad_pvt := make(map[string]float64)
	v_0_pvt := y
	v_1_pvt := b0
	v_2_pvt := b1
	v_3_pvt := x1
	v_4_pvt := b2
	v_5_pvt := y1
	s_0_pvt := multiply_pvt(-1.000000, v_0_pvt)
	s_1_pvt := multiply_pvt(v_2_pvt, v_3_pvt)
	s_2_pvt := add_pvt(v_1_pvt, s_1_pvt)
	s_3_pvt := multiply_pvt(v_4_pvt, v_5_pvt)
	s_4_pvt := add_pvt(s_2_pvt, s_3_pvt)
	s_5_pvt := multiply_pvt(s_0_pvt, s_4_pvt)
	s_6_pvt := exp_pvt(s_5_pvt)
	s_7_pvt := add_pvt(1.000000, s_6_pvt)
	s_8_pvt := log_pvt(s_7_pvt)
	const b_s_8_pvt = 1.0
	b_s_7_pvt := 0.0
	b_s_7_pvt += b_s_8_pvt * (d_log_pvt(s_7_pvt))
	b_s_6_pvt := 0.0
	b_s_6_pvt += b_s_7_pvt * (d_add_pvt(1, 1.000000, s_6_pvt))
	b_s_5_pvt := 0.0
	b_s_5_pvt += b_s_6_pvt * (d_exp_pvt(s_5_pvt))
	b_s_4_pvt := 0.0
	b_s_4_pvt += b_s_5_pvt * (d_multiply_pvt(1, s_0_pvt, s_4_pvt))
	b_s_3_pvt := 0.0
	b_s_3_pvt += b_s_4_pvt * (d_add_pvt(1, s_2_pvt, s_3_pvt))
	b_s_2_pvt := 0.0
	b_s_2_pvt += b_s_4_pvt * (d_add_pvt(0, s_2_pvt, s_3_pvt))
	b_s_1_pvt := 0.0
	b_s_1_pvt += b_s_2_pvt * (d_add_pvt(1, v_1_pvt, s_1_pvt))
	b_s_0_pvt := 0.0
	b_s_0_pvt += b_s_5_pvt * (d_multiply_pvt(0, s_0_pvt, s_4_pvt))
	b_v_5_pvt := 0.0
	b_v_5_pvt += b_s_3_pvt * (d_multiply_pvt(1, v_4_pvt, v_5_pvt))
	grad_pvt["y1"] = b_v_5_pvt
	b_v_4_pvt := 0.0
	b_v_4_pvt += b_s_3_pvt * (d_multiply_pvt(0, v_4_pvt, v_5_pvt))
	grad_pvt["b2"] = b_v_4_pvt
	b_v_3_pvt := 0.0
	b_v_3_pvt += b_s_1_pvt * (d_multiply_pvt(1, v_2_pvt, v_3_pvt))
	grad_pvt["x1"] = b_v_3_pvt
	b_v_2_pvt := 0.0
	b_v_2_pvt += b_s_1_pvt * (d_multiply_pvt(0, v_2_pvt, v_3_pvt))
	grad_pvt["b1"] = b_v_2_pvt
	b_v_1_pvt := 0.0
	b_v_1_pvt += b_s_2_pvt * (d_add_pvt(0, v_1_pvt, s_1_pvt))
	grad_pvt["b0"] = b_v_1_pvt
	b_v_0_pvt := 0.0
	b_v_0_pvt += b_s_0_pvt * (d_multiply_pvt(1, -1.000000, v_0_pvt))
	grad_pvt["y"] = b_v_0_pvt
	return s_8_pvt, grad_pvt
}

// numerically compute the value and gradient of "f := log( 1 + exp(-y * (b0 +  b1 * x1 + b2 * y1)))\n"
func ComputeNumerical(b0, b1, b2, x1, y, y1 float64) (float64, map[string]float64) {
	grad_pvt := make(map[string]float64)
	const delta_pvt = 0.000010
	calc_pvt := func() float64 {
		f := log(1 + exp(-y*(b0+b1*x1+b2*y1)))

		return f
	}
	tmp1_pvt := calc_pvt()
	{
		b0 += delta_pvt
		tmp2_pvt := calc_pvt()
		b0 -= delta_pvt
		grad_pvt["b0"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		b1 += delta_pvt
		tmp2_pvt := calc_pvt()
		b1 -= delta_pvt
		grad_pvt["b1"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		b2 += delta_pvt
		tmp2_pvt := calc_pvt()
		b2 -= delta_pvt
		grad_pvt["b2"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		x1 += delta_pvt
		tmp2_pvt := calc_pvt()
		x1 -= delta_pvt
		grad_pvt["x1"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		y += delta_pvt
		tmp2_pvt := calc_pvt()
		y -= delta_pvt
		grad_pvt["y"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
	}
	{
		y1 += delta_pvt
		tmp2_pvt := calc_pvt()
		y1 -= delta_pvt
		grad_pvt["y1"] = (tmp2_pvt - tmp1_pvt) / delta_pvt
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
