package main

import "testing"

func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
		{
			name:     "single character",
			input:    "a",
			expected: false,
		},
		{
			name:     "two different characters",
			input:    "ab",
			expected: false,
		},
		{
			name:     "two same characters",
			input:    "aa",
			expected: true,
		},
		{
			name:     "no duplicates",
			input:    "abcdef",
			expected: false,
		},
		{
			name:     "has duplicates",
			input:    "abcdefc",
			expected: true,
		},
		{
			name:     "unicode no duplicates",
			input:    "汉字한글",
			expected: false,
		},
		{
			name:     "unicode with duplicates",
			input:    "汉字한글汉",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.input)
			if got != tt.expected {
				t.Errorf("findDuplicate(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
