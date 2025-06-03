package vectors

import "testing"

func TestProject(t *testing.T) {
	a := Vector[float64]{3, 3, 0}
	b := Vector[float64]{1, 0, 0} // x-axis

	proj, err := Project(a, b)
	expected := Vector[float64]{3, 0, 0}

	if err != nil {
		t.Fatal(err)
	}
	for i := range proj {
		if !almostEqual(proj[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, proj)
		}
	}

	// Test with zero vector
	_, err = Project(a, Vector[float64]{0, 0, 0})
	if err == nil {
		t.Error("Expected error when projecting onto zero vector")
	}
}
