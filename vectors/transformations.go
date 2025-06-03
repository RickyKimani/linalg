package vectors

import (
	"errors"
	"math"
)

// Reflect reflects vector v across unit normal vector n.
//
// This function implements the reflection formula: v' = v - 2(v·n)n
// It's commonly used in physics simulations (light reflection, particle collisions)
// and computer graphics.
//
// Parameters:
//   - v: Vector to reflect
//   - n: Normal vector for reflection surface (should be normalized)
//
// Returns:
//   - Vector[float64]: The reflected vector
//   - error: An error if vectors have incompatible dimensions
//
// Note: For correct physical reflection, n should be a unit vector (normalized).
// If not, the reflection will be scaled by the magnitude of n.
func Reflect[T, E int | float64](v Vector[T], n Vector[E]) (Vector[float64], error) {
	// v' = v - 2(v·n)n
	dot, err := Dot(v, n)
	if err != nil {
		return nil, err
	}

	scaled := Scale(2*dot, n)
	return Subtract(v, scaled)
}

// Rotate2D rotates a 2D vector by the specified angle in radians.
//
// This function applies a 2D rotation matrix to the vector:
// [x']   [cos θ  -sin θ] [x]
// [y'] = [sin θ   cos θ] [y]
//
// Parameters:
//   - v: 2D vector to rotate
//   - angle: Rotation angle in radians (positive = counterclockwise)
//
// Returns:
//   - Vector[float64]: The rotated vector
//   - error: An error if the vector is not 2D
//
// Time complexity: O(1) - constant time regardless of vector size
func Rotate2D[T int | float64](v Vector[T], angle float64) (Vector[float64], error) {
	if len(v) != 2 {
		return nil, errors.New("vector must be 2D")
	}

	sin, cos := math.Sin(angle), math.Cos(angle)
	return Vector[float64]{
		cos*float64(v[0]) - sin*float64(v[1]),
		sin*float64(v[0]) + cos*float64(v[1]),
	}, nil
}

// Rotate3D rotates a 3D vector around an arbitrary axis by the specified angle.
//
// This function implements Rodrigues' rotation formula to rotate vector v
// around axis k by angle theta (in radians).
//
// Parameters:
//   - v: 3D vector to rotate
//   - axis: Unit vector representing the rotation axis
//   - angle: Rotation angle in radians (positive = right-hand rule)
//
// Returns:
//   - Vector[float64]: The rotated vector
//   - error: An error if the vector is not 3D or if the axis is not a unit vector
func Rotate3D[T, E int | float64](v Vector[T], axis Vector[E], angle float64) (Vector[float64], error) {
	if len(v) != 3 || len(axis) != 3 {
		return nil, errors.New("both vectors must be 3D")
	}

	// Check if axis is approximately a unit vector
	if !IsUnit(axis) {
		return nil, errors.New("axis must be a unit vector")
	}

	// Rodrigues' rotation formula:
	// v_rot = v*cos(θ) + (k×v)*sin(θ) + k*(k·v)*(1-cos(θ))

	cos, sin := math.Cos(angle), math.Sin(angle)

	// Term 1: v*cos(θ)
	term1 := Scale(cos, v)

	// Term 2: (k×v)*sin(θ)
	cross, err := Cross(axis, v)
	if err != nil {
		return nil, err
	}
	term2 := Scale(sin, cross)

	// Term 3: k*(k·v)*(1-cos(θ))
	dot, _ := Dot(axis, v)
	term3 := Scale(dot*(1-cos), axis)

	result, err := Add(term1, term2)
	if err != nil {
		return nil, err
	}

	return Add(result, term3)
}
