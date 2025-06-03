package vectors

import "errors"

// Normalize returns the unit vector (vector of length 1) in the direction of the input vector.
//
// Normalization is the process of scaling a vector to have a magnitude of 1 while
// preserving its direction. Unit vectors are important in physics, computer graphics,
// and many mathematical applications where only direction matters.
//
// Parameters:
//   - a: Input vector to normalize
//
// Returns:
//   - Vector[float64]: A unit vector in the same direction as the input
//   - error: An error if the input is a zero vector (which has no defined direction)
//     or an empty vector
//
// Formula: â = a/|a| where |a| is the magnitude of vector a
//
// Properties of the returned unit vector:
//   - Has magnitude (length) equal to 1
//   - Points in the same direction as the input vector
//   - Each component is scaled by 1/|a|
//
// Example:
//
//	v := Vector[float64]{3.0, 4.0}
//	unit, _ := Normalize(v)  // Returns [0.6, 0.8]
func Normalize[T int | float64](a Vector[T]) (Vector[float64], error) {
	if IsZero(a) {
		return nil, errors.New("cannot normalize zero vector")
	}
	mag := Magnitude(a)
	if mag == 0 {
		return nil, errors.New("cannot normalize zero vector")
	}
	result := make(Vector[float64], len(a))
	for i, v := range a {
		result[i] = float64(v) / mag
	}
	return result, nil
}
