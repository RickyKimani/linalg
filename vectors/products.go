package vectors

import "errors"

// ScalarProduct calculates the scalar triple product of three vectors (A • (B × C)).
//
// The scalar triple product represents the signed volume of the parallelepiped
// defined by the three vectors. It's also equal to the determinant of the 3×3 matrix
// formed by the three vectors as rows or columns.
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//   - c: Third vector of type Vector[B]
//
// Returns:
//   - float64: Scalar triple product A • (B × C)
//   - error: An error if any vector is not 3D or if operations cannot be performed
//
// Properties:
//   - A • (B × C) = B • (C × A) = C • (A × B) (cyclic permutation)
//   - A • (B × C) = -A • (C × B) (anticommutativity of cross product)
//   - A • (B × C) = 0 if vectors are coplanar
//
// Applications:
//   - Volume calculation of parallelepiped
//   - Testing if three vectors are coplanar
//   - Calculating torque in mechanics
//
// Example:
//
//	v1 := Vector[float64]{1, 0, 0}
//	v2 := Vector[float64]{0, 1, 0}
//	v3 := Vector[float64]{0, 0, 1}
//	vol, _ := ScalarProduct(v1, v2, v3) // Returns 1.0 (unit cube volume)
func ScalarProduct[T, E, B int | float64](a Vector[T], b Vector[E], c Vector[B]) (float64, error) {
	// All vectors must be 3D for cross product
	if len(a) != 3 || len(b) != 3 || len(c) != 3 {
		return 0, errors.New("scalar triple product requires three 3D vectors")
	}

	cross, err := Cross(b, c)
	if err != nil {
		return 0, err
	}
	return Dot(a, cross)
}

// VectorProduct calculates the vector triple product (A × (B × C)).
//
// The vector triple product can be rewritten using the BAC-CAB identity:
// A × (B × C) = B(A•C) - C(A•B)
//
// Parameters:
//   - a: First vector of type Vector[T]
//   - b: Second vector of type Vector[E]
//   - c: Third vector of type Vector[B]
//
// Returns:
//   - Vector[float64]: Vector triple product A × (B × C)
//   - error: An error if any vector is not 3D or if operations cannot be performed
//
// Properties:
//   - A × (B × C) lies in the plane defined by B and C
//   - The result is perpendicular to A × B and A × C
//
// Example:
//
//	v1 := Vector[float64]{1, 0, 0}
//	v2 := Vector[float64]{0, 1, 0}
//	v3 := Vector[float64]{0, 0, 1}
//	result, _ := VectorProduct(v1, v2, v3) // Returns [0, -1, 0]
func VectorProduct[T, E, B int | float64](a Vector[T], b Vector[E], c Vector[B]) (Vector[float64], error) {
	// All vectors must be 3D for cross product
	if len(a) != 3 || len(b) != 3 || len(c) != 3 {
		return nil, errors.New("vector triple product requires three 3D vectors")
	}

	// BAC-CAB identity: A × (B × C) = B(A•C) - C(A•B)
	dotAC, err := Dot(a, c)
	if err != nil {
		return nil, err
	}

	dotAB, err := Dot(a, b)
	if err != nil {
		return nil, err
	}

	scaledB := Scale(dotAC, b)
	scaledC := Scale(dotAB, c)

	return Subtract(scaledB, scaledC)
}
