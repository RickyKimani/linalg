package matrix

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	good := [][]int{
		{1, 2},
		{3, 4},
	}

	bad := [][]int{
		{1, 3, 5},
		{2, 1, 1},
		{1},
	}

	// Test valid matrix
	m1, err1 := NewMatrix(good)
	if err1 != nil {
		t.Errorf("NewMatrix failed with valid input: %v", err1)
	}

	// Expected result should be Matrix[float64], not Matrix[int]
	expected := Matrix[float64]{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	if !reflect.DeepEqual(m1, expected) {
		t.Errorf("NewMatrix returned incorrect matrix: got %v, want %v", m1, expected)
	}

	// Test invalid matrix with inconsistent row lengths
	_, err2 := NewMatrix(bad)
	if err2 == nil {
		t.Error("NewMatrix should fail with inconsistent row lengths")
	}

	// Test deep copy - modifications to original shouldn't affect matrix
	m3, _ := NewMatrix(good)
	good[0][0] = 99       // Modify the original data
	if m3[0][0] == 99.0 { // Compare against 99.0 since m3 is Matrix[float64]
		t.Error("NewMatrix should create a deep copy, not share memory with input")
	}

	// Test empty matrix
	empty := [][]int{}
	m4, err4 := NewMatrix(empty)
	if err4 != nil {
		t.Errorf("NewMatrix failed with empty input: %v", err4)
	}
	if len(m4) != 0 {
		t.Errorf("NewMatrix with empty input should return empty matrix, got %v", m4)
	}

	// Test nil input
	var nilInput [][]int = nil
	m5, err5 := NewMatrix(nilInput)
	if err5 != nil {
		t.Errorf("NewMatrix failed with nil input: %v", err5)
	}
	if len(m5) != 0 {
		t.Errorf("NewMatrix with nil input should return empty matrix, got %v", m5)
	}

	// Test float64 input
	goodFloat := [][]float64{
		{1.1, 2.2},
		{3.3, 4.4},
	}
	m6, err6 := NewMatrix(goodFloat)
	if err6 != nil {
		t.Errorf("NewMatrix failed with float64 input: %v", err6)
	}
	expectedFloat := Matrix[float64]{
		{1.1, 2.2},
		{3.3, 4.4},
	}
	if !reflect.DeepEqual(m6, expectedFloat) {
		t.Errorf("NewMatrix returned incorrect float64 matrix: got %v, want %v", m6, expectedFloat)
	}
}

func TestAbs(t *testing.T) {

	if abs(5) != 5 {
		t.Errorf("abs(5) = %d, want 5", abs(5))
	}

	if abs(-5) != 5 {
		t.Errorf("abs(-5) = %d, want 5", abs(-5))
	}

	if abs(0) != 0 {
		t.Errorf("abs(0) = %d, want 0", abs(0))
	}

	if abs(5.5) != 5.5 {
		t.Errorf("abs(5.5) = %f, want 5.5", abs(5.5))
	}

	if abs(-5.5) != 5.5 {
		t.Errorf("abs(-5.5) = %f, want 5.5", abs(-5.5))
	}

	if abs(0.0) != 0.0 {
		t.Errorf("abs(0.0) = %f, want 0.0", abs(0.0))
	}
}

func TestIsSquare(t *testing.T) {
	emptyMatrix := Matrix[int]{}
	if emptyMatrix.isSquare() {
		t.Error("Empty matrix should not be square")
	}

	square1x1 := Matrix[int]{{1}}
	if !square1x1.isSquare() {
		t.Error("1x1 matrix should be square")
	}

	square2x2 := Matrix[int]{
		{1, 2},
		{3, 4},
	}
	if !square2x2.isSquare() {
		t.Error("2x2 matrix should be square")
	}

	square3x3 := Matrix[int]{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	if !square3x3.isSquare() {
		t.Error("3x3 matrix should be square")
	}

	wideMatrix := Matrix[int]{
		{1, 2, 3},
		{4, 5, 6},
	}
	if wideMatrix.isSquare() {
		t.Error("2x3 matrix should not be square")
	}

	tallMatrix := Matrix[int]{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	if tallMatrix.isSquare() {
		t.Error("3x2 matrix should not be square")
	}
}
