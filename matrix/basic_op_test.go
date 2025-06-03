package matrix

import "testing"

func TestAdd(t *testing.T) {
	A := Matrix[int]{
		{1, 2},
		{3, 4},
	}
	B := Matrix[float64]{
		{0.5, 1.5},
		{2.5, 3.5},
	}

	C := Matrix[int]{}

	D := Matrix[int]{
		{1, 2, 3},
		{4, 5, 6},
	}
	E := Matrix[int]{
		{1},
	}

	got, err := Add(A, B)
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	want := Matrix[float64]{
		{1.5, 3.5},
		{5.5, 7.5},
	}

	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("Add: got[%d][%d]=%v, want=%v", i, j, got[i][j], want[i][j])
			}
		}
	}

	got, err = Add(A, C)
	if err == nil {
		t.Fatalf("Expected error but got %v", got)
	}
	got, err = Add(A, D)
	if err == nil {
		t.Fatalf("Expected error but got %v", got)
	}
	got, err = Add(A, E)
	if err == nil {
		t.Fatalf("Expected error but got %v", got)
	}

}

func TestSubtract(t *testing.T) {
	A := Matrix[int]{
		{5, 6},
		{7, 8},
	}
	B := Matrix[float64]{
		{1.5, 2.5},
		{3.5, 4.5},
	}
	C := Matrix[int]{}

	D := Matrix[int]{
		{1, 2, 3},
		{4, 5, 6},
	}
	E := Matrix[int]{
		{1},
	}

	got, err := Subtract(A, B)
	if err != nil {
		t.Fatalf("Subtract returned error: %v", err)
	}

	want := Matrix[float64]{
		{3.5, 3.5},
		{3.5, 3.5},
	}

	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("Subtract: got[%d][%d]=%v, want=%v", i, j, got[i][j], want[i][j])
			}
		}
	}
	got, err = Subtract(A, C)
	if err == nil {
		t.Fatalf("Expected error but got %v", got)
	}
	got, err = Subtract(A, D)
	if err == nil {
		t.Fatalf("Expected error but got %v", got)
	}
	got, err = Subtract(A, E)
	if err == nil {
		t.Fatalf("Expected error but got %v", got)
	}
}
