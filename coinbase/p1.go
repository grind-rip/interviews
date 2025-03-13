package main

import (
	"strconv"
	"strings"
)

// getTotalClicksByDomain returns a map with the total number of clicks for
// each domain.
func getTotalClicksByDomain(arr []string) map[string]int {
	m := make(map[string]int)
	for _, elem := range arr {
		// Parse the string into substrings.
		// Ex: "900,google.com" -> ["900", "google.com"]
		ss := strings.Split(elem, ",")
		// Get the total number of clicks.
		clicks, _ := strconv.Atoi(ss[0])
		// Get all domain parts.
		// Ex: "google.com" -> ["google", "com"]
		parts := strings.Split(ss[1], ".")
		// Add all domains with corresponding clicks to the hash map.
		for i := len(parts) - 1; i >= 0; i-- {
			subdomain := strings.Join(parts[i:], ".")
			m[subdomain] += clicks
		}
	}
	return m
}
