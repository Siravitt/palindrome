package palindrome

import (
	"strconv"
)

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

func Palindrome(n int) bool {
	s := strconv.Itoa(n)
	if len(s) == 1 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
