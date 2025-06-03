package vectors

import "errors"

// DirectionCosines calculates the direction cosines of a 3D vector.
//
// Direction cosines represent the cosines of the angles between a vector and
// the three coordinate axes (x, y, and z). They effectively describe the
// "contribution" of each axis to the vector's direction.
//
// Parameters:
//   - a: Input 3D vector
//
// Returns:
//   - l: Cosine of the angle between the vector and x-axis
//   - m: Cosine of the angle between the vector and y-axis
//   - n: Cosine of the angle between the vector and z-axis
//   - err: Error if the vector is not 3D or is a zero vector
//
// Direction cosines always satisfy l² + m² + n² = 1, and each component is
// in the range [-1, 1]. They are equivalent to the components of the unit vector
// in the same direction as the input vector.
func DirectionCosines[T int | float64](a Vector[T]) (l, m, n float64, err error) {
	if len(a) != 3 {
		return 0, 0, 0, errors.New("direction cosines are only defined for 3D vectors")
	}

	// Direction cosines are undefined for the zero vector
	if IsZero(a) {
		return 0, 0, 0, errors.New("cannot calculate direction cosines for zero vector")
	}

	// Direction cosines are the components of the unit vector
	unitVector, _ := Normalize(a)

	return unitVector[0], unitVector[1], unitVector[2], nil
}
