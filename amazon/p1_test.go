package main

import (
	"testing"
)

func TestGetMinMoves(t *testing.T) {
	tests := []struct {
		name     string
		plates   []int32
		expected int32
	}{
		{
			name:     "simple reverse ordered array",
			plates:   []int32{3, 2, 1},
			expected: 3,
		},
		{
			name:     "mixed order array",
			plates:   []int32{1, 20, 6, 4, 5},
			expected: 5,
		},
		{
			name:     "longer mixed order array",
			plates:   []int32{10, 3, 4, 2, 5, 7, 9, 11},
			expected: 8,
		},
		{
			name:     "single element array",
			plates:   []int32{1},
			expected: 0,
		},
		{
			name:     "already sorted array",
			plates:   []int32{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "two elements in wrong order",
			plates:   []int32{2, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (naive)", func(t *testing.T) {
			// Create a copy of plates for naive implementation
			platesCopy := make([]int32, len(tt.plates))
			copy(platesCopy, tt.plates)

			got := getMinMovesNaive(platesCopy)
			if got != tt.expected {
				t.Errorf("getMinMovesNaive() = %v, want %v", got, tt.expected)
			}
		})

		t.Run(tt.name+" (optimal)", func(t *testing.T) {
			// Create a copy of plates for optimal implementation
			platesCopy := make([]int32, len(tt.plates))
			copy(platesCopy, tt.plates)

			got := getMinMovesOptimal(platesCopy)
			if got != tt.expected {
				t.Errorf("getMinMovesOptimal() = %v, want %v", got, tt.expected)
			}
		})
	}
}
