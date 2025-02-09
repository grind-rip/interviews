package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "single character",
			s:    "a",
			want: true,
		},
		{
			name: "two same characters",
			s:    "aa",
			want: true,
		},
		{
			name: "two different characters",
			s:    "ab",
			want: false,
		},
		{
			name: "palindrome odd length",
			s:    "level",
			want: true,
		},
		{
			name: "palindrome even length",
			s:    "noon",
			want: true,
		},
		{
			name: "not palindrome",
			s:    "hello",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestIsPalindromeRecursive(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "single character",
			s:    "a",
			want: true,
		},
		{
			name: "two same characters",
			s:    "aa",
			want: true,
		},
		{
			name: "two different characters",
			s:    "ab",
			want: false,
		},
		{
			name: "palindrome odd length",
			s:    "level",
			want: true,
		},
		{
			name: "palindrome even length",
			s:    "noon",
			want: true,
		},
		{
			name: "not palindrome",
			s:    "hello",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeRecursive(tt.s); got != tt.want {
				t.Errorf("isPalindromeRecursive(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
