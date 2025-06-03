package vectors

import (
	"math"
	"testing"
)

func BenchmarkMagnitude(b *testing.B) {
	v := Vector[float64]{1, 2, 3, 4, 5}
	for b.Loop() {
		_ = Magnitude(v)
	}
}

func BenchmarkScale(b *testing.B) {
	v := Vector[float64]{1, 2, 3, 4, 5}
	for b.Loop() {
		_ = Scale(2.5, v)
	}
}

func BenchmarkIsOrthogonal(b *testing.B) {
	a := Vector[float64]{1, 0, 0}
	c := Vector[float64]{0, 1, 0}
	for b.Loop() {
		_, _ = IsOrthogonal(a, c)
	}
}

func BenchmarkProject(b *testing.B) {
	a := Vector[float64]{3, 4, 5}
	c := Vector[float64]{1, 0, 0}
	for b.Loop() {
		_, _ = Project(a, c)
	}
}

func BenchmarkCartesianToSpherical(b *testing.B) {
	v := Vector[float64]{1, 2, 3}
	for b.Loop() {
		_, _, _, _ = CartesianToSpherical(v)
	}
}

func BenchmarkVectorProduct(b *testing.B) {
	a := Vector[float64]{1, 2, 3}
	c := Vector[float64]{4, 5, 6}
	d := Vector[float64]{7, 8, 9}
	for b.Loop() {
		_, _ = VectorProduct(a, c, d)
	}
}

func BenchmarkRotate3D(b *testing.B) {
	v := Vector[float64]{1, 0, 0}
	axis := Vector[float64]{0, 0, 1}
	angle := math.Pi / 4
	for b.Loop() {
		_, _ = Rotate3D(v, axis, angle)
	}
}
