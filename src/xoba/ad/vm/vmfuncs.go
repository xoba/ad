package vm

import "math"

func exp10(f float64) float64 {
	return math.Pow(10, f)
}

func d_add_d0(a, b float64) float64 {
	return 1
}
func d_add_d1(a, b float64) float64 {
	return 1
}

func d_divide_d0(a, b float64) float64 {
	return 1 / b
}
func d_divide_d1(a, b float64) float64 {
	return -a / (b * b)
}

func d_multiply_d0(a, b float64) float64 {
	return b
}
func d_multiply_d1(a, b float64) float64 {
	return a
}

func d_subtract_d0(a, b float64) float64 {
	return 1
}
func d_subtract_d1(a, b float64) float64 {
	return -1
}

func d_pow_d0(a, b float64) float64 {
	return b * math.Pow(a, b-1)
}
func d_pow_d1(a, b float64) float64 {
	return math.Pow(a, b) * math.Log(a)
}

func d_abs_d0(a float64) float64 {
	switch {
	case a > 0:
		return +1
	case a < 0:
		return -1
	default:
		return 0
	}
}
func d_acos_d0(a float64) float64 {
	return -1 / math.Sqrt(1-a*a)
}
func d_asin_d0(a float64) float64 {
	return 1 / math.Sqrt(1-a*a)
}
func d_atan_d0(a float64) float64 {
	return 1 / (1 + a*a)
}
func d_cos_d0(a float64) float64 {
	return -math.Sin(a)
}
func d_cosh_d0(a float64) float64 {
	return math.Sinh(a)
}
func d_exp_d0(a float64) float64 {
	return math.Exp(a)
}
func d_exp10_d0(a float64) float64 {
	panic("d_exp10_d0 unimplemented")
}
func d_exp2_d0(a float64) float64 {
	return math.Log(2) * math.Pow(2, a)
}
func d_log_d0(a float64) float64 {
	return 1 / a
}
func d_log10_d0(a float64) float64 {
	panic("d_log10_d0 unimplemented")
}

var cst_log2 float64 = math.Log(2)

func d_log2_d0(a float64) float64 {
	return 1 / (a * cst_log2)
}
func d_sin_d0(a float64) float64 {
	return math.Cos(a)
}
func d_sinh_d0(a float64) float64 {
	return math.Cosh(a)
}
func d_sqrt_d0(a float64) float64 {
	return 0.5 * math.Pow(a, -0.5)
}
func d_tan_d0(a float64) float64 {
	return math.Pow(1/math.Cos(a), 2)
}
func d_tanh_d0(a float64) float64 {
	return 1 - math.Pow(math.Tanh(a), 2)
}
