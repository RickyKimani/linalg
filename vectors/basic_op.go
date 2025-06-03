package vectors

import "errors"

// Add combines two vectors by element-wise addition.
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//
// Returns:
//   - Vector[float64]: A new vector where each element is the sum of the
//     corresponding elements in a and b
//   - error: An error if vectors have incompatible dimensions
//
// The result is always a float64 vector to accommodate mixed-type operations.
func Add[T, E int | float64](a Vector[T], b Vector[E]) (Vector[float64], error) {
	if len(a) != len(b) {
		return nil, errors.New("vectors must have the same dimension")
	}
	result := make(Vector[float64], len(a))
	for i := range a {
		result[i] = float64(a[i]) + float64(b[i])
	}
	return result, nil
}

// Subtract creates a new vector by element-wise subtraction of the second vector
// from the first vector.
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//
// Returns:
//   - Vector[float64]: A new vector where each element is a[i] - b[i]
//   - error: An error if vectors have incompatible dimensions
//
// The result is always a float64 vector to accommodate mixed-type operations.
func Subtract[T, E int | float64](a Vector[T], b Vector[E]) (Vector[float64], error) {
	if len(a) != len(b) {
		return nil, errors.New("vectors must have the same dimension")
	}
	result := make(Vector[float64], len(a))
	for i := range a {
		result[i] = float64(a[i]) - float64(b[i])
	}
	return result, nil
}

// Negate returns the additive inverse of a vector (reverses all elements).
//
// Parameters:
//   - v: Input vector
//
// Returns:
//   - Vector[float64]: A new vector where each element is the negation of the input
//
// This is equivalent to scaling the vector by -1.
func Negate[T int | float64](v Vector[T]) Vector[float64] {
	return Scale(-1, v)
}
