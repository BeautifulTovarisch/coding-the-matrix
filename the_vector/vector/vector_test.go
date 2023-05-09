// Tests for the Vector package
package vector

import (
	"testing"
)

func mapEq(a, b map[string]float64) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}

	return true
}

func scaleMap(alpha float64, m map[string]float64) map[string]float64 {
	res := make(map[string]float64)

	for k, v := range m {
		res[k] = v * alpha
	}

	return res
}

func TestLookup(t *testing.T) {
	t.Run("Lookup key found", func(t *testing.T) {
		expected := 1
		v := Vector[int]{[]string{"A"}, map[string]int{"A": 1}}

		if val, ok := v.Lookup("A"); val != expected || !ok {
			t.Errorf("Expected %d. Got %d. OK? %t", expected, val, ok)
		}
	})

	t.Run("Lookup key not found", func(t *testing.T) {
		expected := 0
		v := ZeroVec([]string{"A", "B"})

		if val, ok := v.Lookup("A"); val != expected || ok {
			t.Errorf("Expected %d. Got %d. OK? %t", expected, val, ok)
		}
	})
}

// Neglecting complexity of precise floating-point comparision
func TestScalarMul(t *testing.T) {
	keys := []string{"A", "B"}

	t.Run("ScaleMul(1.0)", func(t *testing.T) {
		f := map[string]float64{"A": 1.0, "B": 2.0}

		v := Vector[float64]{keys, f}
		newVec := v.ScalarMul(1.0)

		if actual := newVec.F; !mapEq(v.F, actual) {
			t.Errorf("Scale (%v): Expected %v. Got %v", v.F, v.F, actual)
		}
	})

	t.Run("ScaleMul(2.0)", func(t *testing.T) {
		f := map[string]float64{"A": 1.0, "B": 2.0}

		v := Vector[float64]{keys, f}
		expected := scaleMap(2.0, f)
		newVec := v.ScalarMul(2.0)

		if actual := newVec.F; !mapEq(expected, actual) {
			t.Errorf("Scale (%v): Expected %v. Got %v", v.F, expected, actual)
		}
	})

	t.Run("Scale Zero Vector", func(t *testing.T) {
		// "hack" to perform type conversion of function image
		v := ZeroVec(keys).ScalarMul(1.0)
		nv := v.ScalarMul(1.25)

		if actual := nv.F; !mapEq(v.F, actual) {
			t.Errorf("Expected no-op on Zero Vector. Got %v", actual)
		}
	})
}
