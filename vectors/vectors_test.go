package vectors

import "testing"

func TestVectorsAlmostEqual(t *testing.T) {
	a := Vector[float64]{1, 2, 3}
	b := Vector[float64]{1.0000001, 2.0000001, 3.0000001}
	c := Vector[float64]{1.1, 2.1, 3.1}

	equal, err := vectorsAlmostEqual(a, b)
	if err != nil || !equal {
		t.Errorf("Expected vectors to be equal within epsilon")
	}

	equal, err = vectorsAlmostEqual(a, c)
	if err != nil || equal {
		t.Errorf("Expected vectors to not be equal within epsilon")
	}

	_, err = vectorsAlmostEqual(a, Vector[float64]{1, 2})
	if err == nil {
		t.Error("Expected error for different dimensions")
	}
}

func TestIsZero(t *testing.T) {
	zeroVec := Vector[float64]{0, 0, 0}
	nonZeroVec := Vector[float64]{0, 0, 0.001}

	if !IsZero(zeroVec) {
		t.Error("Expected true for zero vector")
	}

	if IsZero(nonZeroVec) {
		t.Error("Expected false for non-zero vector")
	}
}
func TestIsUnit(t *testing.T) {
	unitVec := Vector[float64]{0, 1, 0}
	nonUnitVec := Vector[float64]{1, 1, 1}

	if !IsUnit(unitVec) {
		t.Error("Expected true for unit vector")
	}

	if IsUnit(nonUnitVec) {
		t.Error("Expected false for non-unit vector")
	}
}
func TestIsOrthogonal(t *testing.T) {
	a := Vector[float64]{1, 0, 0}
	b := Vector[float64]{0, 1, 0}
	c := Vector[float64]{1, 1, 0}

	orthogonal, err := IsOrthogonal(a, b)
	if err != nil || !orthogonal {
		t.Error("Expected vectors to be orthogonal")
	}

	orthogonal, err = IsOrthogonal(a, c)
	if err != nil || orthogonal {
		t.Error("Expected vectors to not be orthogonal")
	}
}

func TestIsParallel(t *testing.T) {
	a := Vector[float64]{1, 0, 0}
	b := Vector[float64]{2, 0, 0}
	c := Vector[float64]{-3, 0, 0}
	d := Vector[float64]{1, 1, 0}

	parallel, err := IsParallel(a, b)
	if err != nil || !parallel {
		t.Error("Expected vectors to be parallel")
	}

	parallel, err = IsParallel(a, c)
	if err != nil || !parallel {
		t.Error("Expected vectors to be parallel (opposite direction)")
	}

	parallel, err = IsParallel(a, d)
	if err != nil || parallel {
		t.Error("Expected vectors to not be parallel")
	}
}
