// Package linalg
package linalg

// NumberLike is a type constraint that permits any signed or unsigned integer
// or floating-point type.
type NumberLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}
