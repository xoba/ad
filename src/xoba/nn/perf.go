package nn

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

func (t PerformanceTracker) P() int {
	return t.tp + t.fn
}
func (t PerformanceTracker) N() int {
	return t.fp + t.tn
}

func (t PerformanceTracker) Accuracy() float64 {
	return float64(t.tp+t.tn) / float64(t.P()+t.N())
}
