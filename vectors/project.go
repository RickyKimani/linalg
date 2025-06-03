package vectors

import "errors"

// Project calculates the vector projection of vector a onto vector b.
//
// Vector projection gives the component of a in the direction of b.
// The result is a vector parallel to b with magnitude |a|·cos(θ),
// where θ is the angle between a and b.
//
// Parameters:
//   - a: Vector to be projected
//   - b: Vector onto which to project
//
// Returns:
//   - Vector[float64]: The projection of a onto b
//   - error: An error if vectors have incompatible dimensions, are empty, or if b is a zero vector
//
// Formula: proj_b(a) = (a·b/|b|²)·b = ((a·b)/(b·b))·b
//
// Properties:
//   - The result is parallel to b
//   - If a and b are perpendicular, the result is the zero vector
//   - If a and b are parallel, the result equals a (if same direction) or -a (if opposite)
//
// Example:
//
//	v1 := Vector[float64]{3.0, 4.0}
//	v2 := Vector[float64]{1.0, 0.0}
//	proj, _ := Project(v1, v2)  // Returns [3.0, 0.0] (projection onto x-axis)
func Project[T, E int | float64](a Vector[T], b Vector[E]) (Vector[float64], error) {
	if len(a) != len(b) {
		return nil, errors.New("vectors must have the same dimension")
	}
	if len(a) == 0 {
		return nil, errors.New("vectors cannot be empty")
	}

	if IsZero(b) {
		return nil, errors.New("cannot project onto zero vector")
	}

	dot, err := Dot(a, b)
	if err != nil {
		return nil, err
	}

	bMagSq := 0.0
	for _, v := range b {
		bMagSq += float64(v) * float64(v)
	}

	scale := dot / bMagSq

	result := make(Vector[float64], len(a))
	for i, v := range b {
		result[i] = scale * float64(v)
	}

	return result, nil
}
