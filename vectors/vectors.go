// Package vectors provides generic operations for mathematical vectors
// supporting both integer and floating-point elements.
//
// The package uses Go's generics to operate on vectors of different
// numeric types while ensuring type safety. It includes functions for
// basic vector operations, geometric calculations, and vector analysis.
package vectors

import (
	"errors"
	"fmt"
	"math"
)

// Vector represents a mathematical vector as a slice of numeric values.
//
// The generic type parameter T is constrained to either int or float64,
// allowing for both integer and floating-point vectors.
type Vector[T int | float64] []T

// NewVector creates a new Vector[float64] from a slice of numeric values.
//
// This function converts all elements from the input slice to float64 type,
// making it useful for creating vectors from different numeric sources while
// ensuring consistent float64 output for mathematical operations.
//
// Parameters:
//   - S: A slice of type T (int or float64) containing vector components
//
// Returns:
//   - Vector[float64]: A new vector with all components converted to float64 type
//
// Example:
//
//	intSlice := []int{1, 2, 3}
//	vec := NewVector(intSlice)  // Returns Vector[float64]{1.0, 2.0, 3.0}
func NewVector[T int | float64](S []T) Vector[float64] {
	var vec Vector[float64]
	for _, val := range S {
		vec = append(vec, float64(val))
	}
	return vec
}

// IsZero checks if a vector has all zero components.
//
// A zero vector is significant in vector spaces as it's the additive identity
// and has special properties like having no defined direction.
//
// Returns:
//   - bool: true if all components are zero, false otherwise
func IsZero[T int | float64](v Vector[T]) bool {
	for _, val := range v {
		if val != 0 {
			return false
		}
	}
	return true
}

// IsUnit checks if a vector has magnitude approximately equal to 1.
//
// Unit vectors are important in many applications as they represent
// pure direction without magnitude. The function uses a small epsilon
// value to account for floating-point precision errors.
//
// Returns:
//   - bool: true if the vector's magnitude is approximately 1.0, false otherwise
func IsUnit[T int | float64](v Vector[T]) bool {
	return math.Abs(Magnitude(v)-1.0) < 1e-10
}

// IsOrthogonal checks if two vectors are perpendicular (orthogonal) to each other.
//
// Two vectors are orthogonal if their dot product is zero. The function uses a small
// epsilon value to account for floating-point precision errors.
//
// Parameters:
//   - a: First vector
//   - b: Second vector
//
// Returns:
//   - bool: true if the vectors are orthogonal, false otherwise
//   - error: An error if the vectors have incompatible dimensions
func IsOrthogonal[T, E int | float64](a Vector[T], b Vector[E]) (bool, error) {
	dot, err := Dot(a, b)
	if err != nil {
		return false, fmt.Errorf("checking orthogonality: %w", err)
	}
	return math.Abs(dot) < 1e-10, nil
}

// IsParallel checks if two vectors are parallel or anti-parallel.
//
// Two vectors are parallel if one is a scalar multiple of the other.
// Parallel vectors have the same or opposite directions.
//
// Parameters:
//   - a: First vector
//   - b: Second vector
//
// Returns:
//   - bool: true if vectors are parallel, false otherwise
//   - error: An error if the vectors have incompatible dimensions or if either is a zero vector
func IsParallel[T, E int | float64](a Vector[T], b Vector[E]) (bool, error) {
	if IsZero(a) || IsZero(b) {
		return false, errors.New("zero vectors have no defined direction")
	}

	if len(a) != len(b) {
		return false, errors.New("vectors must have the same dimension")
	}

	// Normalize both vectors to compare directions
	aNorm, err := Normalize(a)
	if err != nil {
		return false, err
	}

	bNorm, err := Normalize(b)
	if err != nil {
		return false, err
	}

	// Check if normalized vectors are equal or negatives of each other
	dotProduct, _ := Dot(aNorm, bNorm)
	return math.Abs(math.Abs(dotProduct)-1.0) < 1e-10, nil
}

func vectorsAlmostEqual[T, E int | float64](a Vector[T], b Vector[E]) (bool, error) {
	if len(a) != len(b) {
		return false, errors.New("vectors must have the same dimension")
	}

	const epsilon = 1e-6

	for i := range a {
		if math.Abs(float64(a[i])-float64(b[i])) > epsilon {
			return false, nil
		}
	}

	return true, nil
}

func almostEqual(a, b float64, eps float64) bool {
	return math.Abs(a-b) < eps
}
