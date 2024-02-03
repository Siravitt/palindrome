package palindrome

import "testing"

func TestOneChar(t *testing.T) {
	input := 1
	want := true

	res := Palindrome(input)
	if res != want {
		t.Fatalf("input: %v want %v but got %v", input, want, res)
	}
}
