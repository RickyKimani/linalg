package vectors

import "errors"

// Cross calculates the cross product of two 3D vectors.
//
// The cross product a × b produces a vector that is perpendicular to both input vectors.
// The resulting vector's magnitude equals the area of the parallelogram formed by the inputs,
// and its direction follows the right-hand rule.
//
// Properties of the cross product:
//   - a × b = -(b × a) (anti-commutativity)
//   - a × a = 0 (vector crossed with itself is zero)
//   - |a × b| = |a|·|b|·sin(θ) where θ is the angle between vectors
//
// Parameters:
//   - a: First 3D vector
//   - b: Second 3D vector
//
// Returns:
//   - Vector[float64]: The cross product a × b as a 3D vector
//   - error: Error if either vector is not 3D
//
// The cross product is only defined for 3D vectors.
//
// Formula:
//
//	a × b = [a₂b₃-a₃b₂, a₃b₁-a₁b₃, a₁b₂-a₂b₁]
//
// Example:
//
//	v1 := Vector[float64]{1, 0, 0}  // Unit x-axis
//	v2 := Vector[float64]{0, 1, 0}  // Unit y-axis
//	v3, _ := Cross(v1, v2)          // Returns [0, 0, 1] (unit z-axis)
func Cross[T, E int | float64](a Vector[T], b Vector[E]) (Vector[float64], error) {
	if len(a) != 3 || len(b) != 3 {
		return nil, errors.New("cross product is only defined for 3D vectors")
	}

	return Vector[float64]{
		float64(a[1])*float64(b[2]) - float64(a[2])*float64(b[1]),
		float64(a[2])*float64(b[0]) - float64(a[0])*float64(b[2]),
		float64(a[0])*float64(b[1]) - float64(a[1])*float64(b[0]),
	}, nil
}
