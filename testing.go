package sstats

import (
	"math"
	"testing"
)

type statTest struct {
	Init     Statistic
	Values   []float64
	Valid    bool
	Expected float64
}

func runTests(t *testing.T, tests []statTest) {
	for i, test := range tests {
		stat := test.Init
		stat.Push(test.Values...)
		value, valid := stat.Value()

		if valid != test.Valid {
			t.Errorf("[%d] got valid=%t, expected %t", i, valid, test.Valid)
		}

		// Test for equality first, if it fails, check for a margin.
		const tolerance = 1e-8
		if value != test.Expected && !nearlyEqual(value, test.Expected, tolerance) {
			t.Errorf("[%d] got value=%.8f, expected %.8f", i, value, test.Expected)
		}
	}
}

func nearlyEqual(a, b, epsilon float64) bool {
	// Copied and translated straight from
	// http://floating-point-gui.de/errors/comparison/
	diff := math.Abs(a - b)
	const (
		MinNormal float64 = 2.2250738585072014E-308 // IEEE-754 Double Precision
		MaxValue  float64 = math.MaxFloat64
	)
	if a == b {
		// shortcut, handles infinities
		return true
	} else if a == 0 || b == 0 || diff < MinNormal {
		// a or b is zero or both are extremely close to it
		// relative error is less meaningful here
		return diff < (epsilon * MinNormal)
	} else {
		// Relative error
		absA := math.Abs(a)
		absB := math.Abs(b)
		return diff/math.Min(absA+absB, MaxValue) < epsilon
	}
}
