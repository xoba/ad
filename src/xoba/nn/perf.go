package nn

import "math"

type PerformanceTracker struct {
	tp, tn, fp, fn int
}

func (t *PerformanceTracker) Update(score float64, class bool) {
	if score > 0 {
		if class {
			t.tp++
		} else {
			t.fn++
		}
	} else {
		if class {
			t.fp++
		} else {
			t.tn++
		}
	}
}

func (t PerformanceTracker) Tp() int {
	return t.tp
}

func (t PerformanceTracker) Tn() int {
	return t.tn
}

func (t PerformanceTracker) Fp() int {
	return t.fp
}
func (t PerformanceTracker) TypeIErrors() int {
	return t.fp
}
func (t PerformanceTracker) FalseAlarms() int {
	return t.fp
}
func (t PerformanceTracker) TypeIIErrors() int {
	return t.fn
}
func (t PerformanceTracker) Misses() int {
	return t.fn
}

func (t PerformanceTracker) Fn() int {
	return t.fn
}

func (t PerformanceTracker) P() int {
	return t.tp + t.fn
}
func (t PerformanceTracker) N() int {
	return t.fp + t.tn
}

func (t PerformanceTracker) Accuracy() float64 {
	return float64(t.tp+t.tn) / float64(t.P()+t.N())
}

func (t PerformanceTracker) Sensitivity() float64 {
	return float64(t.tp) / float64(t.P())
}
func (t PerformanceTracker) HitRate() float64 {
	return t.Sensitivity()
}
func (t PerformanceTracker) Recall() float64 {
	return t.Sensitivity()
}

func (t PerformanceTracker) Specificity() float64 {
	return float64(t.tn) / float64(t.N())
}
func (t PerformanceTracker) TrueNegativeRate() float64 {
	return t.Specificity()
}

func (t PerformanceTracker) FallOut() float64 {
	return float64(t.fp) / float64(t.N())
}
func (t PerformanceTracker) FalsePositiveRate() float64 {
	return t.FallOut()
}

func (t PerformanceTracker) FalseDiscoveryRate() float64 {
	return float64(t.fp) / float64(t.tp+t.fp)
}

func (t PerformanceTracker) NegativePredictiveValue() float64 {
	return float64(t.tn) / float64(t.tn+t.fn)
}

func (t PerformanceTracker) F1() float64 {
	return float64(2*t.tp) / float64(2*t.tp+t.fp+t.fn)
}

func (t PerformanceTracker) Mcc() float64 {
	return float64(t.tp*t.tn-t.fp*t.fn) / math.Sqrt(float64(t.tp+t.fp)) / math.Sqrt(float64(t.tp+t.fn)) / math.Sqrt(float64(t.tn+t.fp)) / math.Sqrt(float64(t.tn+t.fn))
}
