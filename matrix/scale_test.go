package matrix

import (
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	// Test case 1: Integer matrix with integer scalar
	intMatrix := Matrix[int]{
		{1, 2, 3},
		{4, 5, 6},
	}
	intScalar := 2
	wantIntInt := Matrix[float64]{
		{2, 4, 6},
		{8, 10, 12},
	}
	gotIntInt := Scale(intScalar, intMatrix)
	if !reflect.DeepEqual(gotIntInt, wantIntInt) {
		t.Errorf("Scale(int, Matrix[int]) = %v, want %v", gotIntInt, wantIntInt)
	}

	// Test case 2: Float matrix with float scalar
	floatMatrix := Matrix[float64]{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
	}
	floatScalar := 0.5
	wantFloatFloat := Matrix[float64]{
		{0.5, 1.0, 1.5},
		{2.0, 2.5, 3.0},
	}
	gotFloatFloat := Scale(floatScalar, floatMatrix)
	if !reflect.DeepEqual(gotFloatFloat, wantFloatFloat) {
		t.Errorf("Scale(float64, Matrix[float64]) = %v, want %v", gotFloatFloat, wantFloatFloat)
	}

	// Test case 3: Integer matrix with float scalar
	wantIntFloat := Matrix[float64]{
		{1.5, 3.0, 4.5},
		{6.0, 7.5, 9.0},
	}
	gotIntFloat := Scale(1.5, intMatrix)
	if !reflect.DeepEqual(gotIntFloat, wantIntFloat) {
		t.Errorf("Scale(float64, Matrix[int]) = %v, want %v", gotIntFloat, wantIntFloat)
	}

	// Test case 4: Float matrix with integer scalar
	wantFloatInt := Matrix[float64]{
		{3.0, 6.0, 9.0},
		{12.0, 15.0, 18.0},
	}
	gotFloatInt := Scale(3, floatMatrix)
	if !reflect.DeepEqual(gotFloatInt, wantFloatInt) {
		t.Errorf("Scale(int, Matrix[float64]) = %v, want %v", gotFloatInt, wantFloatInt)
	}

	// Test case 5: Scaling by zero
	wantZero := Matrix[float64]{
		{0, 0, 0},
		{0, 0, 0},
	}
	gotZero := Scale(0, intMatrix)
	if !reflect.DeepEqual(gotZero, wantZero) {
		t.Errorf("Scale(0, Matrix) = %v, want %v", gotZero, wantZero)
	}

	// Test case 6: Scaling by negative value
	wantNegative := Matrix[float64]{
		{-1, -2, -3},
		{-4, -5, -6},
	}
	gotNegative := Scale(-1, intMatrix)
	if !reflect.DeepEqual(gotNegative, wantNegative) {
		t.Errorf("Scale(-1, Matrix) = %v, want %v", gotNegative, wantNegative)
	}

	// Test case 7: Empty matrix
	emptyMatrix := Matrix[int]{}
	wantEmpty := Matrix[float64]{}
	gotEmpty := Scale(2, emptyMatrix)
	if !reflect.DeepEqual(gotEmpty, wantEmpty) {
		t.Errorf("Scale(2, EmptyMatrix) = %v, want %v", gotEmpty, wantEmpty)
	}

	// Test case 8: Matrix with inconsistent rows
	invalidMatrix := Matrix[int]{
		{1, 2, 3},
		{4, 5}, // Shorter row
	}
	wantInvalid := Matrix[float64]{}
	gotInvalid := Scale(2, invalidMatrix)
	if !reflect.DeepEqual(gotInvalid, wantInvalid) {
		t.Errorf("Scale(2, InvalidMatrix) = %v, want %v", gotInvalid, wantInvalid)
	}
}
