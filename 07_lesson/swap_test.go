package _7_lesson

import (
	"testing"
)

func TestSwap(t *testing.T) {
	a := 9
	b := 4
	swap(&a, &b)
	if a != 4 {
		t.Error()
	}
	if b != 9 {
		t.Error()
	}
	swap_bit(&a, &b)
	if a != 9 {
		t.Error()
	}
	if b != 4 {
		t.Error()
	}
}
