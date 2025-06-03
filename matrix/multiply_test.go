package matrix

import (
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	t.Run("int x int", func(t *testing.T) {
		a := Matrix[int]{
			{1, 2},
			{3, 4},
		}
		b := Matrix[int]{
			{2, 0},
			{1, 2},
		}
		expected := Matrix[float64]{
			{4, 4},
			{10, 8},
		}
		result, err := Multiply(a, b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("float x float", func(t *testing.T) {
		a := Matrix[float64]{
			{1.0, 0.5},
			{2.0, 1.0},
		}
		b := Matrix[float64]{
			{3.0, 1.0},
			{1.0, 2.0},
		}
		expected := Matrix[float64]{
			{3.5, 2.0},
			{7.0, 4.0},
		}
		result, err := Multiply(a, b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("int x float", func(t *testing.T) {
		a := Matrix[int]{
			{1, 2},
			{3, 4},
		}
		b := Matrix[float64]{
			{0.5, 1.5},
			{2.0, 1.0},
		}
		expected := Matrix[float64]{
			{4.5, 3.5},
			{9.5, 8.5},
		}
		result, err := Multiply(a, b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("float x int", func(t *testing.T) {
		a := Matrix[float64]{
			{1.5, 2.5},
			{3.0, 1.0},
		}
		b := Matrix[int]{
			{2, 0},
			{1, 3},
		}
		expected := Matrix[float64]{
			{5.5, 7.5},
			{7.0, 3.0},
		}
		result, err := Multiply(a, b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("dimension mismatch", func(t *testing.T) {
		a := Matrix[int]{
			{1, 2},
			{3, 4},
		}
		b := Matrix[int]{
			{1, 2},
		}
		_, err := Multiply(a, b)
		if err == nil {
			t.Error("expected dimension mismatch error")
		}
	})

	t.Run("first matrix empty", func(t *testing.T) {
		a := Matrix[int]{}
		b := Matrix[int]{
			{1, 2},
			{3, 4},
		}
		_, err := Multiply(a, b)
		if err == nil {
			t.Error("expected empty matrix error")
		}
		if err.Error() != "empty matrix" {
			t.Errorf("expected 'empty matrix' error, got: %v", err)
		}
	})

	t.Run("second matrix empty", func(t *testing.T) {
		a := Matrix[int]{
			{1, 2},
			{3, 4},
		}
		b := Matrix[int]{}
		_, err := Multiply(a, b)
		if err == nil {
			t.Error("expected empty matrix error")
		}
		if err.Error() != "empty matrix" {
			t.Errorf("expected 'empty matrix' error, got: %v", err)
		}
	})
}

func BenchmarkMultiply100x100(b *testing.B) {
	a := randomFloatMatrix(100, 100)
	bm := randomFloatMatrix(100, 100)
	for b.Loop() {
		_, _ = Multiply(a, bm)
	}
}
