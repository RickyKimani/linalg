package vectors

import (
	"errors"
	"math"
)

// EuclideanDistance calculates the Euclidean distance between two vectors.
//
// The Euclidean distance is the "straight-line" distance between two points
// in Euclidean space, given by the Pythagorean formula:
// d(a,b) = √[(a₁-b₁)² + (a₂-b₂)² + ... + (aₙ-bₙ)²]
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//
// Returns:
//   - float64: The Euclidean distance between a and b
//   - error: An error if vectors have incompatible dimensions
//
// Time complexity: O(n) where n is the dimension of the vectors.
//
// Example:
//
//	v1 := Vector[int]{0, 0, 0}
//	v2 := Vector[float64]{1.0, 1.0, 1.0}
//	dist, _ := Distance(v1, v2)  // Returns √3 ≈ 1.732
func EuclideanDistance[T, E int | float64](a Vector[T], b Vector[E]) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must have the same dimension")
	}

	var sum float64
	for i := range a {
		diff := float64(a[i]) - float64(b[i])
		sum += diff * diff
	}

	return math.Sqrt(sum), nil
}

// ManhattanDistance calculates the Manhattan (L1 norm) distance between two vectors.
//
// The Manhattan distance represents the sum of absolute differences between coordinates,
// simulating movement along a grid-like path.
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//
// Returns:
//   - float64: The Manhattan distance between a and b
//   - error: An error if vectors have incompatible dimensions
//
// Formula: d(a,b) = |a₁-b₁| + |a₂-b₂| + ... + |aₙ-bₙ|
//
// https://simple.wikipedia.org/wiki/Manhattan_distance
func ManhattanDistance[T, E int | float64](a Vector[T], b Vector[E]) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must have the same dimension")
	}

	var sum float64
	for i := range a {
		diff := math.Abs(float64(a[i]) - float64(b[i]))
		sum += diff
	}

	return sum, nil
}

// ChebyshevDistance calculates the Chebyshev (L∞ norm) distance between two vectors.
//
// The Chebyshev distance is the maximum absolute difference between any dimension,
// representing the minimal number of moves for a king on a chessboard.
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//
// Returns:
//   - float64: The Chebyshev distance between a and b
//   - error: An error if vectors have incompatible dimensions
//
// Formula: d(a,b) = max(|a₁-b₁|, |a₂-b₂|, ..., |aₙ-bₙ|)
//
// https://en.wikipedia.org/wiki/Chebyshev_distance
func ChebyshevDistance[T, E int | float64](a Vector[T], b Vector[E]) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must have the same dimension")
	}

	var maxDiff float64
	for i := range a {
		diff := math.Abs(float64(a[i]) - float64(b[i]))
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff, nil
}
