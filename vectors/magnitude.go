package vectors

import "math"

// Magnitude calculates the Euclidean norm (length) of a vector.
//
// The magnitude of a vector is the square root of the sum of the squares of its
// components, representing the distance from the origin to the point.
//
// Parameters:
//   - a: Input vector of type Vector[T]
//
// Returns:
//   - float64: The magnitude (length) of the vector, always non-negative
//
// For the zero vector, the function returns 0.
// For a unit vector, the function returns 1.
//
// Formula: |a| = √(a₁² + a₂² + ... + aₙ²)
//
// The magnitude is used in normalizing vectors, calculating distances,
// and determining vector properties like unit vectors.
//
// Example:
//
//	v := Vector[float64]{3.0, 4.0}
//	mag := Magnitude(v)  // Returns 5.0 (Pythagorean triple 3-4-5)
func Magnitude[T int | float64](a Vector[T]) float64 {
	var sum float64
	for _, v := range a {
		sum += float64(v) * float64(v)
	}
	return math.Sqrt(sum)
}
