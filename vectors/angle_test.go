package vectors

import (
	"math"
	"testing"
)

func TestAngle(t *testing.T) {
	a := Vector[float64]{1, 0, 0}
	b := Vector[float64]{0, 1, 0}
	c := Vector[float64]{1, 1, 0}

	// Test perpendicular vectors (90 degrees)
	angle, err := Angle(a, b)
	if err != nil || !almostEqual(angle, math.Pi/2, 1e-6) {
		t.Errorf("Expected π/2, got %v", angle)
	}

	// Test 45 degree angle
	angle, err = Angle(a, c)
	if err != nil || !almostEqual(angle, math.Pi/4, 1e-6) {
		t.Errorf("Expected π/4, got %v", angle)
	}

	// Test zero vector error
	_, err = Angle(Vector[float64]{0, 0, 0}, a)
	if err == nil {
		t.Error("Expected error for zero vector")
	}
}

func TestAngleDeg(t *testing.T) {
	a := Vector[float64]{1, 0, 0}
	b := Vector[float64]{0, 1, 0}

	angle, err := AngleDeg(a, b)
	if err != nil || !almostEqual(angle, 90.0, 1e-6) {
		t.Errorf("Expected 90, got %v", angle)
	}
}
