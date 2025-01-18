package main

import (
	"testing"
)

func TestFindMaximumSustainableClusterSize(t *testing.T) {
	tests := []struct {
		name            string
		processingPower []int32
		bootingPower    []int32
		powerMax        int64
		want            int32
	}{
		{
			name:            "example case",
			processingPower: []int32{2, 1, 3, 4, 5},
			bootingPower:    []int32{3, 6, 1, 3, 4},
			powerMax:        25,
			want:            3,
		},
		{
			name:            "single processor",
			processingPower: []int32{3},
			bootingPower:    []int32{5},
			powerMax:        10,
			want:            1,
		},
		{
			name:            "two processors below limit",
			processingPower: []int32{1, 1},
			bootingPower:    []int32{1, 1},
			powerMax:        5,
			want:            2,
		},
		{
			name:            "all processors exceed power limit",
			processingPower: []int32{10, 10, 10},
			bootingPower:    []int32{10, 10, 10},
			powerMax:        5,
			want:            0,
		},
		{
			name:            "large cluster possible",
			processingPower: []int32{1, 1, 1, 1, 1},
			bootingPower:    []int32{1, 1, 1, 1, 1},
			powerMax:        100,
			want:            5,
		},
		{
			name:            "maximum booting power limits cluster size",
			processingPower: []int32{1, 1, 1, 1},
			bootingPower:    []int32{1, 10, 1, 1},
			powerMax:        15,
			want:            2,
		},
		{
			name:            "edge case - just at power limit",
			processingPower: []int32{2, 2},
			bootingPower:    []int32{3, 3},
			powerMax:        11, // 3 + (2 + 2) * 2 = 11
			want:            2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (naive)", func(t *testing.T) {
			got := findMaximumSustainableClusterSizeNaive(
				tt.processingPower,
				tt.bootingPower,
				tt.powerMax,
			)
			if got != tt.want {
				t.Errorf(
					"findMaximumSustainableClusterSizeNaive() = %v, want %v",
					got,
					tt.want,
				)
			}
		})

		t.Run(tt.name+" (optimal)", func(t *testing.T) {
			got := findMaximumSustainableClusterSizeOptimal(
				tt.processingPower,
				tt.bootingPower,
				tt.powerMax,
			)
			if got != tt.want {
				t.Errorf(
					"findMaximumSustainableClusterSizeOptimal() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
