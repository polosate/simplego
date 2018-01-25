package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "pasdjbabobaks"
	if longestPalindrome(s) != "aboba" {
		t.Error()
	}
}

func TestConverDiag(t *testing.T) {
	s := "PAYPALISHIRING"
	expected := "PAHNAPLSIIGYIR"
	n := 3

	s1 := convert(s, n)
	if s1 != expected {
		t.Error("want", expected, "have", s1)
	}
}

