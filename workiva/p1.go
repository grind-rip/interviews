package main

// Given a string `s`, return `true` if it is a palindrome, or `false` otherwise.
//
// NOTES
//   - Use two pointers.

// isPalindrome returns `true` if it is a palindrome, or `false` otherwise.
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// isPalindromeRecursive returns `true` if it is a palindrome, or `false`
// otherwise. Uses a recursive approach instead of iterative.
func isPalindromeRecursive(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindromeRecursive(s[1 : len(s)-1])
}
