package vectors

// Scale multiplies each element of a vector by a scalar value.
//
// Parameters:
//   - scalar: Value to multiply each vector element by
//   - v: Input vector
//
// Returns:
//   - Vector[float64]: A new vector with each element multiplied by the scalar value
//
// The result is always a float64 vector to accommodate mixed-type operations.
func Scale[S, T int | float64](scalar S, v Vector[T]) Vector[float64] {
	result := make(Vector[float64], len(v))
	for i := range v {
		result[i] = float64(scalar) * float64(v[i])
	}
	return result
}
