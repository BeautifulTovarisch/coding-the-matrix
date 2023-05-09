// Package Vector implements a representation of a Vector as function.
package vector

// Type constraint on the basic Numeric types
type Numeric interface {
	uint | int | float32 | float64
}

// Vector represents a vector as a function. The struct exposes a reference to
// its Domain and Function. The Domain is stored as a set, and the public API
// of the vector package assume uniqueness of the keys in D.
type Vector[T Numeric] struct {
	D []string     // D(omain) of the function [F].
	F map[string]T // F(unction)
}

// ZeroVec returns the trivial Vector using a sparse representation. The Vector
// represents the function given by F : [D] -> 0. Thus, the [F] is represented
// by the empty map.
func ZeroVec(domain []string) Vector[int] {
	return Vector[int]{D: domain}
}

// Lookup(key) returns the value of (F(key), true) if the key is present in F
// and the (zero-value of type T, false) otherwise.
//
// Example:
//
//	v := Vector[int]{[]string{"A"}, map[string]int{"A": 1}}
//	v.Lookup("A") // (1, true)
//
//	v := ZeroVec([]string{"A", "B"})
//	v.Lookup("A") // (0, false)
func (v Vector[T]) Lookup(key string) (T, bool) {
	elem, ok := v.F[key]

	return elem, ok
}

// Return new slice rather than modify [elems]
func scale[T Numeric](alpha float64, elems map[string]T) map[string]float64 {
	res := make(map[string]float64)

	for k, v := range elems {
		res[k] = alpha * float64(v)
	}

	return res
}

// ScalarMul(alpha) performs scalar multiplication over the image of F in v. A
// NEW vector is returned with the same domain rather than modifying [v].
//
// The type T of v must be Numeric or the function will panic.
//
// Example:
//
//	v := Vector[int]([]string{"A", "B"}, map[string]int{"A": 2, "B": 3})
//	v.ScalarMul(2) // {"A": 4, "B": 6}
func (v Vector[T]) ScalarMul(alpha float64) Vector[float64] {
	return Vector[float64]{D: v.D, F: scale(alpha, v.F)}
}
