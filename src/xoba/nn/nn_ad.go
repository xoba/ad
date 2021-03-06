// autogenerated, do not edit!
// see https://github.com/xoba/ad

package nn

import (
	"math"
)

// automatically compute the value and gradient of "f := log2( 1 + exp2(-z * (beta[20] +  beta[0] * (1 / (1 + exp2(- (beta[1] + beta[2] * x1 + beta[3] * x2)))) + beta[4] * (1 / (1 + exp2(- (beta[5] + beta[6] * x1 + beta[7] * x2)))) + beta[8] * (1 / (1 + exp2(- (beta[9] + beta[10] * x1 + beta[11] * x2)))) + beta[12] * (1 / (1 + exp2(- (beta[13] + beta[14] * x1 + beta[15] * x2)))) + beta[16] * (1 / (1 + exp2(- (beta[17] + beta[18] * x1 + beta[19] * x2)))))))\n"
func ComputeAD(computeGrad bool, x1, x2, z float64, beta []float64) (float64 /* value */, float64 /* d/d(x1) */, float64 /* d/d(x2) */, float64 /* d/d(z) */, []float64 /* d/d(beta) */) {
	v_0_pvt := z
	v_1_pvt := beta[20]
	v_2_pvt := beta[0]
	v_3_pvt := beta[1]
	v_4_pvt := beta[2]
	v_5_pvt := x1
	v_6_pvt := beta[3]
	v_7_pvt := x2
	v_8_pvt := beta[4]
	v_9_pvt := beta[5]
	v_10_pvt := beta[6]
	v_11_pvt := beta[7]
	v_12_pvt := beta[8]
	v_13_pvt := beta[9]
	v_14_pvt := beta[10]
	v_15_pvt := beta[11]
	v_16_pvt := beta[12]
	v_17_pvt := beta[13]
	v_18_pvt := beta[14]
	v_19_pvt := beta[15]
	v_20_pvt := beta[16]
	v_21_pvt := beta[17]
	v_22_pvt := beta[18]
	v_23_pvt := beta[19]
	s_0_pvt := multiply_pvt(-1.000000, v_0_pvt)
	s_1_pvt := multiply_pvt(v_4_pvt, v_5_pvt)
	s_2_pvt := add_pvt(v_3_pvt, s_1_pvt)
	s_3_pvt := multiply_pvt(v_6_pvt, v_7_pvt)
	s_4_pvt := add_pvt(s_2_pvt, s_3_pvt)
	s_5_pvt := multiply_pvt(-1.000000, s_4_pvt)
	s_6_pvt := exp2_pvt(s_5_pvt)
	s_7_pvt := add_pvt(1.000000, s_6_pvt)
	s_8_pvt := divide_pvt(1.000000, s_7_pvt)
	s_9_pvt := multiply_pvt(v_2_pvt, s_8_pvt)
	s_10_pvt := add_pvt(v_1_pvt, s_9_pvt)
	s_11_pvt := multiply_pvt(v_10_pvt, v_5_pvt)
	s_12_pvt := add_pvt(v_9_pvt, s_11_pvt)
	s_13_pvt := multiply_pvt(v_11_pvt, v_7_pvt)
	s_14_pvt := add_pvt(s_12_pvt, s_13_pvt)
	s_15_pvt := multiply_pvt(-1.000000, s_14_pvt)
	s_16_pvt := exp2_pvt(s_15_pvt)
	s_17_pvt := add_pvt(1.000000, s_16_pvt)
	s_18_pvt := divide_pvt(1.000000, s_17_pvt)
	s_19_pvt := multiply_pvt(v_8_pvt, s_18_pvt)
	s_20_pvt := add_pvt(s_10_pvt, s_19_pvt)
	s_21_pvt := multiply_pvt(v_14_pvt, v_5_pvt)
	s_22_pvt := add_pvt(v_13_pvt, s_21_pvt)
	s_23_pvt := multiply_pvt(v_15_pvt, v_7_pvt)
	s_24_pvt := add_pvt(s_22_pvt, s_23_pvt)
	s_25_pvt := multiply_pvt(-1.000000, s_24_pvt)
	s_26_pvt := exp2_pvt(s_25_pvt)
	s_27_pvt := add_pvt(1.000000, s_26_pvt)
	s_28_pvt := divide_pvt(1.000000, s_27_pvt)
	s_29_pvt := multiply_pvt(v_12_pvt, s_28_pvt)
	s_30_pvt := add_pvt(s_20_pvt, s_29_pvt)
	s_31_pvt := multiply_pvt(v_18_pvt, v_5_pvt)
	s_32_pvt := add_pvt(v_17_pvt, s_31_pvt)
	s_33_pvt := multiply_pvt(v_19_pvt, v_7_pvt)
	s_34_pvt := add_pvt(s_32_pvt, s_33_pvt)
	s_35_pvt := multiply_pvt(-1.000000, s_34_pvt)
	s_36_pvt := exp2_pvt(s_35_pvt)
	s_37_pvt := add_pvt(1.000000, s_36_pvt)
	s_38_pvt := divide_pvt(1.000000, s_37_pvt)
	s_39_pvt := multiply_pvt(v_16_pvt, s_38_pvt)
	s_40_pvt := add_pvt(s_30_pvt, s_39_pvt)
	s_41_pvt := multiply_pvt(v_22_pvt, v_5_pvt)
	s_42_pvt := add_pvt(v_21_pvt, s_41_pvt)
	s_43_pvt := multiply_pvt(v_23_pvt, v_7_pvt)
	s_44_pvt := add_pvt(s_42_pvt, s_43_pvt)
	s_45_pvt := multiply_pvt(-1.000000, s_44_pvt)
	s_46_pvt := exp2_pvt(s_45_pvt)
	s_47_pvt := add_pvt(1.000000, s_46_pvt)
	s_48_pvt := divide_pvt(1.000000, s_47_pvt)
	s_49_pvt := multiply_pvt(v_20_pvt, s_48_pvt)
	s_50_pvt := add_pvt(s_40_pvt, s_49_pvt)
	s_51_pvt := multiply_pvt(s_0_pvt, s_50_pvt)
	s_52_pvt := exp2_pvt(s_51_pvt)
	s_53_pvt := add_pvt(1.000000, s_52_pvt)
	s_54_pvt := log2_pvt(s_53_pvt)
	if !computeGrad {
		return s_54_pvt, 0, 0, 0, nil
	}
	var b_pvt_x1 float64
	var b_pvt_x2 float64
	var b_pvt_z float64
	b_pvt_beta := make([]float64, 21)
	const b_s_54_pvt = 1.0
	b_s_53_pvt := 0.0
	b_s_53_pvt += b_s_54_pvt * (d_log2_pvt(s_53_pvt))
	b_s_52_pvt := 0.0
	b_s_52_pvt += b_s_53_pvt * (d_add_pvt(1, 1.000000, s_52_pvt))
	b_s_51_pvt := 0.0
	b_s_51_pvt += b_s_52_pvt * (d_exp2_pvt(s_51_pvt))
	b_s_50_pvt := 0.0
	b_s_50_pvt += b_s_51_pvt * (d_multiply_pvt(1, s_0_pvt, s_50_pvt))
	b_s_49_pvt := 0.0
	b_s_49_pvt += b_s_50_pvt * (d_add_pvt(1, s_40_pvt, s_49_pvt))
	b_s_48_pvt := 0.0
	b_s_48_pvt += b_s_49_pvt * (d_multiply_pvt(1, v_20_pvt, s_48_pvt))
	b_s_47_pvt := 0.0
	b_s_47_pvt += b_s_48_pvt * (d_divide_pvt(1, 1.000000, s_47_pvt))
	b_s_46_pvt := 0.0
	b_s_46_pvt += b_s_47_pvt * (d_add_pvt(1, 1.000000, s_46_pvt))
	b_s_45_pvt := 0.0
	b_s_45_pvt += b_s_46_pvt * (d_exp2_pvt(s_45_pvt))
	b_s_44_pvt := 0.0
	b_s_44_pvt += b_s_45_pvt * (d_multiply_pvt(1, -1.000000, s_44_pvt))
	b_s_43_pvt := 0.0
	b_s_43_pvt += b_s_44_pvt * (d_add_pvt(1, s_42_pvt, s_43_pvt))
	b_s_42_pvt := 0.0
	b_s_42_pvt += b_s_44_pvt * (d_add_pvt(0, s_42_pvt, s_43_pvt))
	b_s_41_pvt := 0.0
	b_s_41_pvt += b_s_42_pvt * (d_add_pvt(1, v_21_pvt, s_41_pvt))
	b_s_40_pvt := 0.0
	b_s_40_pvt += b_s_50_pvt * (d_add_pvt(0, s_40_pvt, s_49_pvt))
	b_s_39_pvt := 0.0
	b_s_39_pvt += b_s_40_pvt * (d_add_pvt(1, s_30_pvt, s_39_pvt))
	b_s_38_pvt := 0.0
	b_s_38_pvt += b_s_39_pvt * (d_multiply_pvt(1, v_16_pvt, s_38_pvt))
	b_s_37_pvt := 0.0
	b_s_37_pvt += b_s_38_pvt * (d_divide_pvt(1, 1.000000, s_37_pvt))
	b_s_36_pvt := 0.0
	b_s_36_pvt += b_s_37_pvt * (d_add_pvt(1, 1.000000, s_36_pvt))
	b_s_35_pvt := 0.0
	b_s_35_pvt += b_s_36_pvt * (d_exp2_pvt(s_35_pvt))
	b_s_34_pvt := 0.0
	b_s_34_pvt += b_s_35_pvt * (d_multiply_pvt(1, -1.000000, s_34_pvt))
	b_s_33_pvt := 0.0
	b_s_33_pvt += b_s_34_pvt * (d_add_pvt(1, s_32_pvt, s_33_pvt))
	b_s_32_pvt := 0.0
	b_s_32_pvt += b_s_34_pvt * (d_add_pvt(0, s_32_pvt, s_33_pvt))
	b_s_31_pvt := 0.0
	b_s_31_pvt += b_s_32_pvt * (d_add_pvt(1, v_17_pvt, s_31_pvt))
	b_s_30_pvt := 0.0
	b_s_30_pvt += b_s_40_pvt * (d_add_pvt(0, s_30_pvt, s_39_pvt))
	b_s_29_pvt := 0.0
	b_s_29_pvt += b_s_30_pvt * (d_add_pvt(1, s_20_pvt, s_29_pvt))
	b_s_28_pvt := 0.0
	b_s_28_pvt += b_s_29_pvt * (d_multiply_pvt(1, v_12_pvt, s_28_pvt))
	b_s_27_pvt := 0.0
	b_s_27_pvt += b_s_28_pvt * (d_divide_pvt(1, 1.000000, s_27_pvt))
	b_s_26_pvt := 0.0
	b_s_26_pvt += b_s_27_pvt * (d_add_pvt(1, 1.000000, s_26_pvt))
	b_s_25_pvt := 0.0
	b_s_25_pvt += b_s_26_pvt * (d_exp2_pvt(s_25_pvt))
	b_s_24_pvt := 0.0
	b_s_24_pvt += b_s_25_pvt * (d_multiply_pvt(1, -1.000000, s_24_pvt))
	b_s_23_pvt := 0.0
	b_s_23_pvt += b_s_24_pvt * (d_add_pvt(1, s_22_pvt, s_23_pvt))
	b_s_22_pvt := 0.0
	b_s_22_pvt += b_s_24_pvt * (d_add_pvt(0, s_22_pvt, s_23_pvt))
	b_s_21_pvt := 0.0
	b_s_21_pvt += b_s_22_pvt * (d_add_pvt(1, v_13_pvt, s_21_pvt))
	b_s_20_pvt := 0.0
	b_s_20_pvt += b_s_30_pvt * (d_add_pvt(0, s_20_pvt, s_29_pvt))
	b_s_19_pvt := 0.0
	b_s_19_pvt += b_s_20_pvt * (d_add_pvt(1, s_10_pvt, s_19_pvt))
	b_s_18_pvt := 0.0
	b_s_18_pvt += b_s_19_pvt * (d_multiply_pvt(1, v_8_pvt, s_18_pvt))
	b_s_17_pvt := 0.0
	b_s_17_pvt += b_s_18_pvt * (d_divide_pvt(1, 1.000000, s_17_pvt))
	b_s_16_pvt := 0.0
	b_s_16_pvt += b_s_17_pvt * (d_add_pvt(1, 1.000000, s_16_pvt))
	b_s_15_pvt := 0.0
	b_s_15_pvt += b_s_16_pvt * (d_exp2_pvt(s_15_pvt))
	b_s_14_pvt := 0.0
	b_s_14_pvt += b_s_15_pvt * (d_multiply_pvt(1, -1.000000, s_14_pvt))
	b_s_13_pvt := 0.0
	b_s_13_pvt += b_s_14_pvt * (d_add_pvt(1, s_12_pvt, s_13_pvt))
	b_s_12_pvt := 0.0
	b_s_12_pvt += b_s_14_pvt * (d_add_pvt(0, s_12_pvt, s_13_pvt))
	b_s_11_pvt := 0.0
	b_s_11_pvt += b_s_12_pvt * (d_add_pvt(1, v_9_pvt, s_11_pvt))
	b_s_10_pvt := 0.0
	b_s_10_pvt += b_s_20_pvt * (d_add_pvt(0, s_10_pvt, s_19_pvt))
	b_s_9_pvt := 0.0
	b_s_9_pvt += b_s_10_pvt * (d_add_pvt(1, v_1_pvt, s_9_pvt))
	b_s_8_pvt := 0.0
	b_s_8_pvt += b_s_9_pvt * (d_multiply_pvt(1, v_2_pvt, s_8_pvt))
	b_s_7_pvt := 0.0
	b_s_7_pvt += b_s_8_pvt * (d_divide_pvt(1, 1.000000, s_7_pvt))
	b_s_6_pvt := 0.0
	b_s_6_pvt += b_s_7_pvt * (d_add_pvt(1, 1.000000, s_6_pvt))
	b_s_5_pvt := 0.0
	b_s_5_pvt += b_s_6_pvt * (d_exp2_pvt(s_5_pvt))
	b_s_4_pvt := 0.0
	b_s_4_pvt += b_s_5_pvt * (d_multiply_pvt(1, -1.000000, s_4_pvt))
	b_s_3_pvt := 0.0
	b_s_3_pvt += b_s_4_pvt * (d_add_pvt(1, s_2_pvt, s_3_pvt))
	b_s_2_pvt := 0.0
	b_s_2_pvt += b_s_4_pvt * (d_add_pvt(0, s_2_pvt, s_3_pvt))
	b_s_1_pvt := 0.0
	b_s_1_pvt += b_s_2_pvt * (d_add_pvt(1, v_3_pvt, s_1_pvt))
	b_s_0_pvt := 0.0
	b_s_0_pvt += b_s_51_pvt * (d_multiply_pvt(0, s_0_pvt, s_50_pvt))
	b_v_23_pvt := 0.0
	b_v_23_pvt += b_s_43_pvt * (d_multiply_pvt(0, v_23_pvt, v_7_pvt))
	b_pvt_beta[19] = b_v_23_pvt
	b_v_22_pvt := 0.0
	b_v_22_pvt += b_s_41_pvt * (d_multiply_pvt(0, v_22_pvt, v_5_pvt))
	b_pvt_beta[18] = b_v_22_pvt
	b_v_21_pvt := 0.0
	b_v_21_pvt += b_s_42_pvt * (d_add_pvt(0, v_21_pvt, s_41_pvt))
	b_pvt_beta[17] = b_v_21_pvt
	b_v_20_pvt := 0.0
	b_v_20_pvt += b_s_49_pvt * (d_multiply_pvt(0, v_20_pvt, s_48_pvt))
	b_pvt_beta[16] = b_v_20_pvt
	b_v_19_pvt := 0.0
	b_v_19_pvt += b_s_33_pvt * (d_multiply_pvt(0, v_19_pvt, v_7_pvt))
	b_pvt_beta[15] = b_v_19_pvt
	b_v_18_pvt := 0.0
	b_v_18_pvt += b_s_31_pvt * (d_multiply_pvt(0, v_18_pvt, v_5_pvt))
	b_pvt_beta[14] = b_v_18_pvt
	b_v_17_pvt := 0.0
	b_v_17_pvt += b_s_32_pvt * (d_add_pvt(0, v_17_pvt, s_31_pvt))
	b_pvt_beta[13] = b_v_17_pvt
	b_v_16_pvt := 0.0
	b_v_16_pvt += b_s_39_pvt * (d_multiply_pvt(0, v_16_pvt, s_38_pvt))
	b_pvt_beta[12] = b_v_16_pvt
	b_v_15_pvt := 0.0
	b_v_15_pvt += b_s_23_pvt * (d_multiply_pvt(0, v_15_pvt, v_7_pvt))
	b_pvt_beta[11] = b_v_15_pvt
	b_v_14_pvt := 0.0
	b_v_14_pvt += b_s_21_pvt * (d_multiply_pvt(0, v_14_pvt, v_5_pvt))
	b_pvt_beta[10] = b_v_14_pvt
	b_v_13_pvt := 0.0
	b_v_13_pvt += b_s_22_pvt * (d_add_pvt(0, v_13_pvt, s_21_pvt))
	b_pvt_beta[9] = b_v_13_pvt
	b_v_12_pvt := 0.0
	b_v_12_pvt += b_s_29_pvt * (d_multiply_pvt(0, v_12_pvt, s_28_pvt))
	b_pvt_beta[8] = b_v_12_pvt
	b_v_11_pvt := 0.0
	b_v_11_pvt += b_s_13_pvt * (d_multiply_pvt(0, v_11_pvt, v_7_pvt))
	b_pvt_beta[7] = b_v_11_pvt
	b_v_10_pvt := 0.0
	b_v_10_pvt += b_s_11_pvt * (d_multiply_pvt(0, v_10_pvt, v_5_pvt))
	b_pvt_beta[6] = b_v_10_pvt
	b_v_9_pvt := 0.0
	b_v_9_pvt += b_s_12_pvt * (d_add_pvt(0, v_9_pvt, s_11_pvt))
	b_pvt_beta[5] = b_v_9_pvt
	b_v_8_pvt := 0.0
	b_v_8_pvt += b_s_19_pvt * (d_multiply_pvt(0, v_8_pvt, s_18_pvt))
	b_pvt_beta[4] = b_v_8_pvt
	b_v_7_pvt := 0.0
	b_v_7_pvt += b_s_3_pvt * (d_multiply_pvt(1, v_6_pvt, v_7_pvt))
	b_v_7_pvt += b_s_13_pvt * (d_multiply_pvt(1, v_11_pvt, v_7_pvt))
	b_v_7_pvt += b_s_23_pvt * (d_multiply_pvt(1, v_15_pvt, v_7_pvt))
	b_v_7_pvt += b_s_33_pvt * (d_multiply_pvt(1, v_19_pvt, v_7_pvt))
	b_v_7_pvt += b_s_43_pvt * (d_multiply_pvt(1, v_23_pvt, v_7_pvt))
	b_pvt_x2 = b_v_7_pvt
	b_v_6_pvt := 0.0
	b_v_6_pvt += b_s_3_pvt * (d_multiply_pvt(0, v_6_pvt, v_7_pvt))
	b_pvt_beta[3] = b_v_6_pvt
	b_v_5_pvt := 0.0
	b_v_5_pvt += b_s_1_pvt * (d_multiply_pvt(1, v_4_pvt, v_5_pvt))
	b_v_5_pvt += b_s_11_pvt * (d_multiply_pvt(1, v_10_pvt, v_5_pvt))
	b_v_5_pvt += b_s_21_pvt * (d_multiply_pvt(1, v_14_pvt, v_5_pvt))
	b_v_5_pvt += b_s_31_pvt * (d_multiply_pvt(1, v_18_pvt, v_5_pvt))
	b_v_5_pvt += b_s_41_pvt * (d_multiply_pvt(1, v_22_pvt, v_5_pvt))
	b_pvt_x1 = b_v_5_pvt
	b_v_4_pvt := 0.0
	b_v_4_pvt += b_s_1_pvt * (d_multiply_pvt(0, v_4_pvt, v_5_pvt))
	b_pvt_beta[2] = b_v_4_pvt
	b_v_3_pvt := 0.0
	b_v_3_pvt += b_s_2_pvt * (d_add_pvt(0, v_3_pvt, s_1_pvt))
	b_pvt_beta[1] = b_v_3_pvt
	b_v_2_pvt := 0.0
	b_v_2_pvt += b_s_9_pvt * (d_multiply_pvt(0, v_2_pvt, s_8_pvt))
	b_pvt_beta[0] = b_v_2_pvt
	b_v_1_pvt := 0.0
	b_v_1_pvt += b_s_10_pvt * (d_add_pvt(0, v_1_pvt, s_9_pvt))
	b_pvt_beta[20] = b_v_1_pvt
	b_v_0_pvt := 0.0
	b_v_0_pvt += b_s_0_pvt * (d_multiply_pvt(1, -1.000000, v_0_pvt))
	b_pvt_z = b_v_0_pvt
	return s_54_pvt, b_pvt_x1, b_pvt_x2, b_pvt_z, b_pvt_beta
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
	if i < 0 || i > 1 {
		panic("illegal index")
	}
	return 1
}

func d_add_pvt(i int, a, b float64) float64 {
	if i < 0 || i > 1 {
		panic("illegal index")
	}
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

func exp2(a float64) float64 {
	return math.Exp2(a)
}

func exp2_pvt(a float64) float64 {
	return math.Exp2(a)
}

func d_exp2(a float64) float64 {
	return math.Log(2) * math.Pow(2, a)
}

func d_exp2_pvt(a float64) float64 {
	return math.Log(2) * math.Pow(2, a)
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

func log2(a float64) float64 {
	return math.Log2(a)
}

func log2_pvt(a float64) float64 {
	return math.Log2(a)
}

func d_log2(a float64) float64 {
	return 1 / (a * math.Log(2))
}

func d_log2_pvt(a float64) float64 {
	return 1 / (a * math.Log(2))
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
