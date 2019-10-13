package goblas

import (
	"math"
	"testing"
)

func testpanics(f func(), name string, t *testing.T) {
	b := panics(f)
	if !b {
		t.Errorf("%v should panic and does not", name)
	}
}

// returns true if the function panics
func panics(f func()) (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			b = true
		}
	}()
	f()
	return
}

func dTolEqual(a, b float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	if a == b {
		return true
	}
	m := math.Max(math.Abs(a), math.Abs(b))
	if m > 1 {
		a /= m
		b /= m
	}
	if math.Abs(a-b) < 1e-14 {
		return true
	}
	return false
}

// EqualWithinAbs returns true if a and b have an absolute
// difference of less than tol.
func EqualWithinAbs(a, b, tol float64) bool {
	return a == b || math.Abs(a-b) <= tol
}

// HasNaN returns true if the slice s has any values that are NaN and false
// otherwise.
func HasNaN(s []float64) bool {
	for _, v := range s {
		if math.IsNaN(v) {
			return true
		}
	}
	return false
}
