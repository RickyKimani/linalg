package vectors

import "testing"

func TestScalarProduct(t *testing.T) {
	a := Vector[float64]{1, 0, 0}
	b := Vector[float64]{0, 1, 0}
	c := Vector[float64]{0, 0, 1}

	product, err := ScalarProduct(a, b, c)
	if err != nil || !almostEqual(product, 1.0, 1e-6) {
		t.Errorf("Expected 1.0, got %v", product)
	}

	// Test coplanar vectors
	d := Vector[float64]{1, 1, 0}
	product, err = ScalarProduct(a, b, d)
	if err != nil || !almostEqual(product, 0.0, 1e-6) {
		t.Errorf("Expected 0.0 for coplanar vectors, got %v", product)
	}
}

func TestVectorProduct(t *testing.T) {
	a := Vector[float64]{1, 0, 0}
	b := Vector[float64]{0, 1, 0}
	c := Vector[float64]{0, 0, 1}

	product, err := VectorProduct(a, b, c)
	expected := Vector[float64]{0, 0, 0}

	if err != nil {
		t.Fatal(err)
	}
	for i := range product {
		if !almostEqual(product[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, product)
		}
	}
}
