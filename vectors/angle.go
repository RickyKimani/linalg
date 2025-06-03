package vectors

import (
	"errors"
	"math"
)

// angleCosine calculates the cosine of the angle between two vectors.
//
// This helper function computes the value of cos(θ) = (a·b)/(|a|·|b|) where θ is
// the angle between vectors a and b. The function performs validation checks
// and handles potential numerical issues.
//
// Parameters:
//   - a: First vector
//   - b: Second vector
//
// Returns:
//   - float64: Cosine of the angle between the vectors
//   - error: An error if the vectors have incompatible dimensions, are empty, or if
//     either vector is a zero vector (which has no defined direction)
func angleCosine[T, E int | float64](a Vector[T], b Vector[E]) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must have same dimension")
	}
	if len(a) == 0 {
		return 0, errors.New("vectors cannot be empty")
	}

	// Consider using the IsZero function instead of manual checks
	if IsZero(a) || IsZero(b) {
		return 0, errors.New("cannot compute angle with zero vector")
	}

	// Could also use Dot and Magnitude functions to avoid duplicating logic:
	dot, err := Dot(a, b)
	if err != nil {
		return 0, err
	}

	aMag := Magnitude(a)
	bMag := Magnitude(b)

	cos := dot / (aMag * bMag)

	// Clamp to [-1, 1] to handle floating point errors
	const epsilon = 1e-10
	if math.Abs(cos) > 1.0 && math.Abs(cos-1.0) < epsilon {
		return math.Copysign(1.0, cos), nil
	}

	return math.Max(-1.0, math.Min(1.0, cos)), nil
}

// Angle calculates the angle (in radians) between two vectors.
//
// The angle between two vectors is the minimum rotation required to align their
// directions, ranging from 0 (parallel) to π (anti-parallel).
//
// Parameters:
//   - a: First vector
//   - b: Second vector
//
// Returns:
//   - float64: Angle in radians between the vectors (0 to π)
//   - error: An error if the vectors have incompatible dimensions, are empty, or if
//     either vector is a zero vector
func Angle[T, E int | float64](a Vector[T], b Vector[E]) (float64, error) {
	cos, err := angleCosine(a, b)
	if err != nil {
		return 0, err
	}
	return math.Acos(cos), nil
}

// AngleDeg calculates the angle (in degrees) between two vectors.
//
// This is the same as Angle() but returns the result in degrees
// rather than radians for convenience.
//
// Parameters:
//   - a: First vector
//   - b: Second vector
//
// Returns:
//   - float64: Angle in degrees between the vectors (0 to 180)
//   - error: An error if the vectors have incompatible dimensions, are empty, or if
//     either vector is a zero vector
func AngleDeg[T, E int | float64](a Vector[T], b Vector[E]) (float64, error) {
	cos, err := angleCosine(a, b)
	if err != nil {
		return 0, err
	}
	return 180 * math.Acos(cos) / math.Pi, nil
}
