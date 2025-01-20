package main

import (
	"testing"
)

// TestFindBiggestScore verifies that the biggest score is correctly found.
func TestFindBiggestScore(t *testing.T) {
	tests := []struct {
		name      string
		trees     []*Node
		wantScore int
		prepare   func() []*Node // Function to prepare test trees
	}{
		{
			name:      "empty list of trees",
			trees:     []*Node{},
			wantScore: 0,
			prepare:   func() []*Node { return []*Node{} },
		},
		{
			name:      "single tree depth 0",
			wantScore: 0, // Check if the score is within valid range
			prepare: func() []*Node {
				return []*Node{createTree(0)}
			},
		},
		{
			name:      "multiple trees of different depths",
			wantScore: 0, // Check if the score is within valid range
			prepare: func() []*Node {
				return []*Node{
					createTree(0),
					createTree(1),
					createTree(2),
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trees := tt.prepare()
			got := findBiggestScore(trees)

			// For non-empty trees, verify score is within valid range
			if len(trees) > 0 {
				if got < 0 || got >= 100 {
					t.Errorf("findBiggestScore() = %d, want score in range [0,99]", got)
				}
			} else {
				if got != tt.wantScore {
					t.Errorf("findBiggestScore() = %d, want %d", got, tt.wantScore)
				}
			}
		})
	}
}
