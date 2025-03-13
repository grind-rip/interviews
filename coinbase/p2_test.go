package main

import (
	"reflect"
	"testing"
)

func TestFindContiguousHistory(t *testing.T) {
	tests := []struct {
		name     string
		user1    []string
		user2    []string
		expected []string
	}{
		{
			name:     "Basic example with common substring",
			user1:    []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"},
			user2:    []string{"/start", "/pink", "/register", "/orange", "/red", "a"},
			expected: []string{"/pink", "/register", "/orange"},
		},
		{
			name:     "No common substring",
			user1:    []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"},
			user2:    []string{"a", "/one", "/two"},
			expected: []string{},
		},
		{
			name:     "Same user histories",
			user1:    []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"},
			user2:    []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"},
			expected: []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"},
		},
		{
			name:     "Single character match",
			user1:    []string{"a", "/one", "/two"},
			user2:    []string{"/start", "/pink", "/register", "/orange", "/red", "a"},
			expected: []string{"a"},
		},
		{
			name:     "Single element history match",
			user1:    []string{"a"},
			user2:    []string{"a", "/one", "/two"},
			expected: []string{"a"},
		},
		{
			name:     "Longer example 1",
			user1:    []string{"/pink", "/orange", "/yellow", "/plum", "/blue", "/tan", "/red", "/amber", "/HotRodPink", "/CornflowerBlue", "/LightGoldenRodYellow", "/BritishRacingGreen"},
			user2:    []string{"/pink", "/orange", "/amber", "/BritishRacingGreen", "/plum", "/blue", "/tan", "/red", "/lavender", "/HotRodPink", "/CornflowerBlue", "/LightGoldenRodYellow"},
			expected: []string{"/plum", "/blue", "/tan", "/red"},
		},
		{
			name:     "Longer example 2",
			user1:    []string{"/pink", "/orange", "/yellow", "/plum", "/blue", "/tan", "/red", "/amber", "/HotRodPink", "/CornflowerBlue", "/LightGoldenRodYellow", "/BritishRacingGreen"},
			user2:    []string{"/pink", "/orange", "/six", "/plum", "/seven", "/tan", "/red", "/amber"},
			expected: []string{"/tan", "/red", "/amber"},
		},
		{
			name:     "Empty first user history",
			user1:    []string{},
			user2:    []string{"/start", "/pink", "/register"},
			expected: []string{},
		},
		{
			name:     "Empty second user history",
			user1:    []string{"/start", "/pink", "/register"},
			user2:    []string{},
			expected: []string{},
		},
		{
			name:     "Both empty user histories",
			user1:    []string{},
			user2:    []string{},
			expected: []string{},
		},
		{
			name:     "Multiple possible common substrings with same length",
			user1:    []string{"/a", "/b", "/c", "/d", "/e", "/a", "/b"},
			user2:    []string{"/z", "/a", "/b", "/y", "/x", "/a", "/b"},
			expected: []string{"/a", "/b"}, // First occurrence should be returned
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findContiguousHistory(tt.user1, tt.user2)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("findContiguousHistory(%v, %v) = %v, want %v",
					tt.user1, tt.user2, got, tt.expected)
			}
		})
	}
}
