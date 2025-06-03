package vectors

import (
	"math"
	"testing"
)

func TestReflect(t *testing.T) {
	v := Vector[float64]{1, 1, 0}
	normal := Vector[float64]{0, 1, 0} // y-axis as normal

	reflected, err := Reflect(v, normal)
	expected := Vector[float64]{1, -1, 0}

	if err != nil {
		t.Fatal(err)
	}
	for i := range reflected {
		if !almostEqual(reflected[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, reflected)
		}
	}
}

func TestRotate2D(t *testing.T) {
	v := Vector[float64]{1, 0}
	angle := math.Pi / 2 // 90 degrees

	rotated, err := Rotate2D(v, angle)
	expected := Vector[float64]{0, 1}

	if err != nil {
		t.Fatal(err)
	}
	for i := range rotated {
		if !almostEqual(rotated[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, rotated)
		}
	}
}

func TestRotate3D(t *testing.T) {
	v := Vector[float64]{1, 0, 0}
	axis := Vector[float64]{0, 0, 1} // z-axis
	angle := math.Pi / 2             // 90 degrees

	rotated, err := Rotate3D(v, axis, angle)
	expected := Vector[float64]{0, 1, 0}

	if err != nil {
		t.Fatal(err)
	}
	for i := range rotated {
		if !almostEqual(rotated[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, rotated)
		}
	}
}
