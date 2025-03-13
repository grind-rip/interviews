package main

// findContiguousHistory returns the longest contiguous sequence of URLs
// that appears in both s1 and s2.
func findContiguousHistory(s1 []string, s2 []string) []string {
	// Allocate a zeroed 2D array of len(s1) x len(s2).
	m, n := len(s1), len(s2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// Maintain the length of the longest contiguous sequence of URLs and the
	// index of its last element.
	longest, longestIdx := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				// NOTE: Since the solution to the current subproblem builds from the
				// solution to the previous subproblem, it is common to initialize the
				// 2D array to be m + 1 and n + 1, then iterate from dp[1...m][1...n].
				// This, however, requires that the indices into `s1` and `s2` be
				// different from the indices into dp for `i` and `j`. Instead, we can
				// simply handle the edge case with a conditional and set the value
				// accordingly.
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if dp[i][j] > longest {
					longest, longestIdx = dp[i][j], i
				}
			}
		}
	}
	// If longest != 0, return a slice of `s1` using the index of the last
	// element and the length of the longest contiguous sequence.
	if longest == 0 {
		return make([]string, 0)
	} else {
		return s1[longestIdx-longest+1 : longestIdx+1]
	}
}
