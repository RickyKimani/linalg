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

// Origin3D represents the zero vector in 3D space
var Origin3D = Vector[float64]{0, 0, 0}

// UnitX represents the unit vector along the x-axis in 3D space
var UnitX = Vector[float64]{1, 0, 0}

// UnitY represents the unit vector along the y-axis in 3D space
var UnitY = Vector[float64]{0, 1, 0}

// UnitZ represents the unit vector along the z-axis in 3D space
var UnitZ = Vector[float64]{0, 0, 1}

// StandardBasis returns the complete set of standard basis vectors for the given dimension.
//
// Parameters:
//   - dim: The dimension of the vector space
//
// Returns:
//   - []Vector[float64]: A slice containing all basis vectors, essentially an identity matrix of order dim
//   - error: Error if dimension is invalid
//
// Example:
//
//	basis, _ := StandardBasis(3) // Returns [[1,0,0], [0,1,0], [0,0,1]]
func StandardBasis(dim int) ([]Vector[float64], error) {
	if dim <= 0 {
		return nil, errors.New("dimension must be positive")
	}
	basis := make([]Vector[float64], dim)
	for i := range dim {
		basis[i] = make(Vector[float64], dim)
		basis[i][i] = 1
	}
	return basis, nil
}

// StandardBasisVector returns a unit vector for the specified axis and dimension.
//
// Parameters:
//   - axis: The axis index (0 for x, 1 for y, etc.)
//   - dim: The total dimension of the vector space
//
// Returns:
//   - Vector[float64]: A unit vector along the specified axis
//   - error: Error if axis is out of bounds or dimension is invalid
//
// Example:
//
//	xAxis, _ := StandardBasisVector(0, 3) // Returns [1,0,0]
//	yAxis, _ := StandardBasisVector(1, 4) // Returns [0,1,0,0] in 4D space
func StandardBasisVector(axis, dim int) (Vector[float64], error) {
	if dim <= 0 {
		return nil, errors.New("dimension must be positive")
	}

	if axis < 0 || axis >= dim {
		return nil, fmt.Errorf("axis %d out of bounds for dimension %d", axis, dim)
	}

	v := make(Vector[float64], dim)
	v[axis] = 1.0
	return v, nil
}

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
	vec := make(Vector[float64], len(S))
	for i, val := range S {
		vec[i] = float64(val)
	}
	return vec
}

// NewEmptyVector creates a new zero vector of the specified size.
//
// This function allocates a new vector with all components initialized to zero.
// It's useful for creating vectors that will be populated later or for
// initializing accumulator vectors in mathematical operations.
//
// Parameters:
//   - size: The number of components in the vector
//
// Returns:
//   - Vector[float64]: A new zero vector with the specified size
//
// Example:
//
//	vec := NewEmptyVector(3)  // Returns Vector[float64]{0.0, 0.0, 0.0}
func NewEmptyVector(size int) Vector[float64] {
	if size <= 0 {
		return Vector[float64]{}
	}
	return make(Vector[float64], size)
}

// Set modifies the value at the specified index in the vector.
//
// This method provides bounds-checked access for setting vector components.
// It ensures that the index is valid before attempting to modify the vector,
// preventing potential runtime panics.
//
// Parameters:
//   - i: The index of the component to modify (0-based)
//   - val: The new value to set at the specified index
//
// Returns:
//   - error: An error if the index is out of bounds
//
// Example:
//
//	var vec Vector[float64] = Vector[float64]{1.0, 2.0, 3.0}
//	err := vec.Set(1, 5.0)  // vec becomes {1.0, 5.0, 3.0}
func (v *Vector[T]) Set(i int, val T) error {
	if i < 0 || i >= len(*v) {
		return fmt.Errorf("index %d out of bounds for vector of length %d", i, len(*v))
	}
	(*v)[i] = val
	return nil
}

// Get retrieves the value at the specified index in the vector.
//
// This method provides bounds-checked access for reading vector components.
// It ensures that the index is valid before attempting to access the vector,
// preventing potential runtime panics.
//
// Parameters:
//   - i: The index of the component to retrieve (0-based)
//
// Returns:
//   - T: The value at the specified index
//   - error: An error if the index is out of bounds
//
// Example:
//
//	vec := Vector[float64]{1.0, 2.0, 3.0}
//	val, err := vec.Get(1)  // Returns 2.0, nil
func (v *Vector[T]) Get(i int) (T, error) {
	var zero T
	if i < 0 || i >= len(*v) {
		return zero, fmt.Errorf("index %d out of bounds for vector of length %d", i, len(*v))
	}
	return (*v)[i], nil
}

// Size returns the number of components in the vector.
//
// This method provides a convenient way to get the dimensionality of a vector,
// which is equivalent to calling len() on the underlying slice but offers
// better semantic clarity in mathematical contexts.
//
// Returns:
//   - int: The number of components (dimension) of the vector
//
// Example:
//
//	vec := Vector[float64]{1.0, 2.0, 3.0}
//	dim := vec.Size()  // Returns 3
func (v *Vector[T]) Size() int {
	return len(*v)
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
		return false, fmt.Errorf("vectors must have same dimension: got %d and %d", len(a), len(b))
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

// Copy creates a deep copy of the vector, converting it to Vector[float64].
//
// This method creates a new vector with all elements copied and converted
// to float64 type. The original vector remains unchanged, making this useful
// for creating independent copies that can be modified without affecting
// the original.
//
// Returns:
//   - Vector[float64]: A new vector containing copies of all elements as float64
//
// Example:
//
//	original := Vector[int]{1, 2, 3}
//	copy := original.Copy()  // Returns Vector[float64]{1.0, 2.0, 3.0}
//	copy.Set(0, 99.0)        // original remains unchanged
func (v Vector[T]) Copy() Vector[float64] {
	if len(v) == 0 {
		return Vector[float64]{}
	}

	result := make(Vector[float64], len(v))
	for i, val := range v {
		result[i] = float64(val)
	}
	return result
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
