package vectors

import (
	"math"
	"testing"
)

func TestCartesianToPolar(t *testing.T) {
	v := Vector[float64]{3, 4}
	r, theta, err := CartesianToPolar(v)

	if err != nil || !almostEqual(r, 5.0, 1e-6) || !almostEqual(theta, math.Atan2(4, 3), 1e-6) {
		t.Errorf("Expected r=5, theta=atan2(4,3), got r=%v, theta=%v", r, theta)
	}
}

func TestPolarToCartesian(t *testing.T) {
	r := 5.0
	theta := math.Pi / 4 // 45 degrees

	v, err := PolarToCartesian(r, theta)
	expected := Vector[float64]{5 * math.Cos(math.Pi/4), 5 * math.Sin(math.Pi/4)}

	if err != nil {
		t.Fatal(err)
	}
	for i := range v {
		if !almostEqual(v[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	}
}

func TestCartesianToSpherical(t *testing.T) {
	v := Vector[float64]{1, 1, 1}
	rho, theta, phi, err := CartesianToSpherical(v)

	expectedRho := math.Sqrt(3)
	expectedTheta := math.Pi / 4
	expectedPhi := math.Acos(1 / math.Sqrt(3))

	if err != nil || !almostEqual(rho, expectedRho, 1e-6) ||
		!almostEqual(theta, expectedTheta, 1e-6) ||
		!almostEqual(phi, expectedPhi, 1e-6) {
		t.Errorf("Unexpected spherical coordinates: rho=%v, theta=%v, phi=%v", rho, theta, phi)
	}
}
