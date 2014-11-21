package templates

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestOneDims(t *testing.T) {
	test1d("tan", tan, d_tan, t)
	test1d("abs", abs, d_abs, t)
	test1d("atan", atan, d_atan, t)
	test1d("tanh", tanh, d_tanh, t)
	test1d("sin", sin, d_sin, t)
	test1d("asin", asin, d_asin, t)
	test1d("sinh", sinh, d_sinh, t)
	test1d("cos", cos, d_cos, t)
	test1d("acos", acos, d_acos, t)
	test1d("cosh", cosh, d_cosh, t)
	test1d("sqrt", sqrt, d_sqrt, t)
	test1d("exp", exp, d_exp, t)
	test1d("exp2", exp2, d_exp2, t)
	test1d("log", log, d_log, t)
	test1d("log2", log2, d_log2, t)
}

func TestTwoDims(t *testing.T) {
	test2d("add", add, d_add, t)
	test2d("multiply", multiply, d_multiply, t)
	test2d("subtract", subtract, d_subtract, t)
	test2d("divide", divide, d_divide, t)
	test2d("pow", pow, d_pow, t)
}

func test2d(name string, f Function2D, df DFunction2D, t *testing.T) {
	var n int
	for n < 10 {
		x := rand.NormFloat64()
		y := rand.NormFloat64()
		v := f(x, y)
		if math.IsNaN(v) || math.IsInf(v, 0) {
			continue
		}
		n++
		for i := 0; i < 2; i++ {
			if df := math.Abs(df(i, x, y) - f.Derivative()(i, x, y)); df > 0.0001 {
				t.Fatalf("oops: df(%d,%f,%f) = %f for %s", i, x, y, df, name)
			}
		}
	}
}

func test1d(name string, f, d Function, t *testing.T) {
	var n int
	for n < 10 {
		x := rand.NormFloat64()
		y := f(x)
		if math.IsNaN(y) || math.IsInf(y, 0) {
			continue
		}
		n++
		if df := math.Abs(d(x) - f.Derivative()(x)); df > 0.0001 {
			t.Fatalf("oops: df(%f) = %f for %s", x, df, name)
		}
	}
}

type Function2D func(a, b float64) float64
type DFunction2D func(i int, a, b float64) float64

type Function func(x float64) float64

const dx = 0.00000001

func (f Function) Derivative() Function {
	return func(x float64) float64 {
		return (f(x+dx) - f(x)) / dx
	}
}

func (f Function2D) Derivative() DFunction2D {
	return func(i int, a, b float64) float64 {
		var g Function
		var x float64
		switch i {
		case 0:
			g = func(x float64) float64 {
				return f(x, b)
			}
			x = a
		case 1:
			g = func(x float64) float64 {
				return f(a, x)
			}
			x = b
		default:
			panic(fmt.Sprintf("illegal index %d", i))
		}
		return g.Derivative()(x)
	}
}
