package main

import (
	"testing"
)

func TestGetTotalClicksByDomain(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name: "Example from problem statement",
			input: []string{
				"900,google.com",
				"60,mail.yahoo.com",
				"10,mobile.sports.yahoo.com",
				"40,sports.yahoo.com",
				"300,yahoo.com",
				"10,stackoverflow.com",
				"20,overflow.com",
				"5,com.com",
				"2,en.wikipedia.org",
				"1,m.wikipedia.org",
				"1,mobile.sports",
				"1,google.co.uk",
			},
			expected: map[string]int{
				"co.uk":                   1,
				"com":                     1345,
				"com.com":                 5,
				"en.wikipedia.org":        2,
				"google.co.uk":            1,
				"google.com":              900,
				"m.wikipedia.org":         1,
				"mail.yahoo.com":          60,
				"mobile.sports":           1,
				"mobile.sports.yahoo.com": 10,
				"org":                     3,
				"overflow.com":            20,
				"sports":                  1,
				"sports.yahoo.com":        50,
				"stackoverflow.com":       10,
				"uk":                      1,
				"wikipedia.org":           3,
				"yahoo.com":               410,
			},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTotalClicksByDomain(tt.input)

			// Check if got and expected maps have the same length
			if len(got) != len(tt.expected) {
				t.Errorf("Maps have different lengths: got %d, want %d", len(got), len(tt.expected))
			}

			// Check each key-value pair
			for domain, expectedClicks := range tt.expected {
				actualClicks, exists := got[domain]
				if !exists {
					t.Errorf("Domain %q missing from got", domain)
					continue
				}
				if actualClicks != expectedClicks {
					t.Errorf("For domain %q: got %d clicks, want %d clicks", domain, actualClicks, expectedClicks)
				}
			}

			// Check for extra domains
			for domain := range got {
				if _, exists := tt.expected[domain]; !exists {
					t.Errorf("Extra domain in got: %q", domain)
				}
			}
		})
	}
}
