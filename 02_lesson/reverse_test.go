package _2_lesson

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s)
	expected := []int{5, 4, 3, 2, 1, 0}
	if !reflect.DeepEqual(s, expected) {
		t.Error("want", expected, "have", s)
	}
}

func TestLeftShift(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	leftShift(s, 2)
	expected := []int{2, 3, 4, 5, 0, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Error("want", expected, "have", s)
	}
}

func TestRightShift(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	rightShift(s, 2)
	expected := []int{4, 5, 0, 1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Error("want", expected, "have", s)
	}
}
