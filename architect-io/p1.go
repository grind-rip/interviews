package main

// findDuplicate returns true if `s` has a duplicate. Otherwise, returns false.
func findDuplicate(s string) bool {
	table := make(map[rune]bool)
	// Loop over `s`. If a character (rune) is in hash set, return true. If not,
	// add the character to the hash set.
	for _, elem := range s {
		_, ok := table[elem]
		if ok {
			return true
		} else {
			table[elem] = true
		}
	}
	return false
}
