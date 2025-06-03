package vectors

import "errors"

// Dot calculates the dot product (scalar product) of two vectors.
//
// The dot product a·b is defined as the sum of the products of the corresponding
// components: a·b = a₁b₁ + a₂b₂ + ... + aₙbₙ.
//
// Properties of dot product:
//   - a·b = b·a (commutativity)
//   - a·b = |a|·|b|·cos(θ) where θ is the angle between vectors
//   - a·b = 0 if vectors are orthogonal (perpendicular)
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//
// Returns:
//   - float64: The dot product of the two vectors
//   - error: An error if vectors have incompatible dimensions or are empty
//
// The dot product is used to calculate work done by a force, to find
// vector projections, and to test orthogonality between vectors.
//
// Example:
//
//	v1 := Vector[int]{1, 2, 3}
//	v2 := Vector[float64]{4.0, 5.0, 6.0}
//	result, _ := Dot(v1, v2)  // Returns 32.0 = 1*4 + 2*5 + 3*6
func Dot[T, E int | float64](a Vector[T], b Vector[E]) (result float64, err error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must have the same dimension")
	}
	if len(a) == 0 {
		return 0, errors.New("vectors cannot be empty")
	}

	for i := range a {
		result += float64(a[i]) * float64(b[i])
	}
	return result, nil
}
