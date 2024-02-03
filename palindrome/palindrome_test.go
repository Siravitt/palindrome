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

func TestRepeatChar(t *testing.T) {
	input := 11
	want := true

	res := Palindrome(input)
	if res != want {
		t.Fatalf("input: %v want %v but got %v", input, want, res)
	}
}

func TestNonRepeatChar(t *testing.T) {
	input := 9999
	want := true

	res := Palindrome(input)
	if res != want {
		t.Fatalf("input: %v want %v but got %v", input, want, res)
	}
}
