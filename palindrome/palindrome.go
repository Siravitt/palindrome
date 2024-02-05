package palindrome

import (
	"math"
	"strconv"
)

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

func Palindrome(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < int(math.Floor(float64(len(s)/2))); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
