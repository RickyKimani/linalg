package matrix

import (
	"fmt"
	"testing"
)

func TestTrace(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected any
		wantErr  bool
		errorMsg string
	}{
		// Integer matrix tests
		{
			name: "3x3 integer matrix",
			input: Matrix[int]{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 15, // 1 + 5 + 9
			wantErr:  false,
		},
		{
			name: "2x2 integer identity matrix",
			input: Matrix[int]{
				{1, 0},
				{0, 1},
			},
			expected: 2, // 1 + 1
			wantErr:  false,
		},
		{
			name: "1x1 integer matrix",
			input: Matrix[int]{
				{42},
			},
			expected: 42,
			wantErr:  false,
		},
		{
			name: "3x3 integer matrix with zeros",
			input: Matrix[int]{
				{0, 1, 2},
				{3, 0, 5},
				{6, 7, 0},
			},
			expected: 0, // 0 + 0 + 0
			wantErr:  false,
		},
		{
			name: "3x3 integer matrix with negative values",
			input: Matrix[int]{
				{-1, 2, 3},
				{4, -5, 6},
				{7, 8, -9},
			},
			expected: -15, // -1 + (-5) + (-9)
			wantErr:  false,
		},

		// Float64 matrix tests
		{
			name: "2x2 float64 matrix",
			input: Matrix[float64]{
				{1.5, 2.5},
				{3.5, 4.5},
			},
			expected: 6.0, // 1.5 + 4.5
			wantErr:  false,
		},
		{
			name: "3x3 float64 matrix with decimals",
			input: Matrix[float64]{
				{0.1, 0.2, 0.3},
				{0.4, 0.5, 0.6},
				{0.7, 0.8, 0.9},
			},
			expected: 1.5, // 0.1 + 0.5 + 0.9
			wantErr:  false,
		},
		{
			name: "1x1 float64 matrix",
			input: Matrix[float64]{
				{3.14159},
			},
			expected: 3.14159,
			wantErr:  false,
		},

		// Error cases
		{
			name: "non-square integer matrix (2x3)",
			input: Matrix[int]{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: 0,
			wantErr:  true,
			errorMsg: "cannot find trace of a non-square matrix",
		},
		{
			name: "non-square integer matrix (3x2)",
			input: Matrix[int]{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			expected: 0,
			wantErr:  true,
			errorMsg: "cannot find trace of a non-square matrix",
		},
		{
			name: "non-square float64 matrix",
			input: Matrix[float64]{
				{1.0, 2.0, 3.0, 4.0},
				{5.0, 6.0, 7.0, 8.0},
			},
			expected: 0.0,
			wantErr:  true,
			errorMsg: "cannot find trace of a non-square matrix",
		},
		{
			name:     "empty integer matrix",
			input:    Matrix[int]{},
			expected: 0,
			wantErr:  true,
			errorMsg: "cannot find trace of an empty matrix",
		},
		{
			name:     "empty float64 matrix",
			input:    Matrix[float64]{},
			expected: 0.0,
			wantErr:  true,
			errorMsg: "cannot find trace of an empty matrix",
		},
		{
			name: "single row matrix (1xN)",
			input: Matrix[int]{
				{1, 2, 3, 4, 5},
			},
			expected: 0,
			wantErr:  true,
			errorMsg: "cannot find trace of a non-square matrix",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test integer matrices
			if intMatrix, ok := tt.input.(Matrix[int]); ok {
				result, err := Trace(intMatrix)

				if tt.wantErr {
					if err == nil {
						t.Errorf("Trace() expected error, got nil")
						return
					}
					if tt.errorMsg != "" && err.Error() != tt.errorMsg {
						t.Errorf("Trace() error = %v, expected %v", err.Error(), tt.errorMsg)
					}
					if result != tt.expected.(int) {
						t.Errorf("Trace() result = %v, expected %v for error case", result, tt.expected)
					}
					return
				}

				if err != nil {
					t.Errorf("Trace() unexpected error = %v", err)
					return
				}

				if result != tt.expected.(int) {
					t.Errorf("Trace() = %v, expected %v", result, tt.expected)
				}
			}

			// Test float64 matrices
			if floatMatrix, ok := tt.input.(Matrix[float64]); ok {
				result, err := Trace(floatMatrix)

				if tt.wantErr {
					if err == nil {
						t.Errorf("Trace() expected error, got nil")
						return
					}
					if tt.errorMsg != "" && err.Error() != tt.errorMsg {
						t.Errorf("Trace() error = %v, expected %v", err.Error(), tt.errorMsg)
					}
					if result != tt.expected.(float64) {
						t.Errorf("Trace() result = %v, expected %v for error case", result, tt.expected)
					}
					return
				}

				if err != nil {
					t.Errorf("Trace() unexpected error = %v", err)
					return
				}

				expected := tt.expected.(float64)
				if abs(result-expected) > 1e-10 {
					t.Errorf("Trace() = %v, expected %v", result, expected)
				}
			}
		})
	}
}

// BenchmarkTrace runs parametrized benchmarks on the Trace function with different matrix types and sizes
func BenchmarkTrace(b *testing.B) {
	sizes := []int{10, 50, 100, 500}

	// Test integer matrices of different sizes
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Int%dx%d", size, size), func(b *testing.B) {
			m := randomIntMatrix(size, size)
			b.ReportAllocs() // Track memory allocations
			b.ResetTimer()   // Don't include setup time

			for b.Loop() {
				_, _ = Trace(m)
			}
		})
	}

	// Test float64 matrices of different sizes
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Float%dx%d", size, size), func(b *testing.B) {
			m := randomFloatMatrix(size, size)
			b.ReportAllocs()
			b.ResetTimer()

			for b.Loop() {
				_, _ = Trace(m)
			}
		})
	}
}
