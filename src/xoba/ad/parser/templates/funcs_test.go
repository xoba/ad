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

func test1d(name string, f, df Function, t *testing.T) {
	df2 := f.Derivative()
	var n, failed int
	for n < runs {
		a := rand.NormFloat64()
		if y := f(a); math.IsNaN(y) || math.IsInf(y, 0) {
			continue
		}
		n++
		if df := math.Abs(df(a) - df2(a)); df > equalityThreshold {
			failed++
		}
	}
	eval(name, n, failed, t)
}

func test2d(name string, f Function2D, df DFunction2D, t *testing.T) {
	df2 := f.Derivative()
	var n, failed int
	for n < runs {
		a, b := rand.NormFloat64(), rand.NormFloat64()
		if y := f(a, b); math.IsNaN(y) || math.IsInf(y, 0) {
			continue
		}
		n++
		for i := 0; i < 2; i++ {
			if df := math.Abs(df(i, a, b) - df2(i, a, b)); df > equalityThreshold {
				failed++
			}
		}
	}
	eval(name, n, failed, t)
}

func eval(name string, n, failed int, t *testing.T) {
	if float64(failed)/float64(n) > failureThreshold {
		t.Fatalf("oops: failed %d / %d for %s\n", failed, n, name)
	}
}

type Function func(x float64) float64
type Function2D func(a, b float64) float64
type DFunction2D func(i int, a, b float64) float64

const (
	failureThreshold  = 0.03
	runs              = 10000
	dx                = 0.000000001
	equalityThreshold = 0.0001
)

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
