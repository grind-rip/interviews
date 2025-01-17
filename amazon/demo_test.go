package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		input    int32
		expected string
	}{
		{
			name:     "invalid input - zero",
			input:    0,
			expected: "",
		},
		{
			name:     "invalid input - negative",
			input:    -1,
			expected: "",
		},
		{
			name:     "invalid input - too large",
			input:    200000,
			expected: "",
		},
		{
			name:     "single digit",
			input:    1,
			expected: "1\n",
		},
		{
			name:     "multiple of 3",
			input:    3,
			expected: "1\n2\nFizz\n",
		},
		{
			name:     "multiple of 5",
			input:    5,
			expected: "1\n2\nFizz\n4\nBuzz\n",
		},
		{
			name:     "multiple of both 3 and 5",
			input:    15,
			expected: "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				FizzBuzz(tt.input)
			})

			if output != tt.expected {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.input, output, tt.expected)
			}
		})
	}
}

// captureOutput captures data written to stdout and returns it as a string.
func captureOutput(f func()) string {
	// Create a pipe to capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function that produces output
	f()

	// Restore stdout and read captured output
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
