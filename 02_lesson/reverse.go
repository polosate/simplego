package _2_lesson

import "strconv"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func leftShift(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rightShift(s []int, n int) {
	reverse(s)
	reverse(s[:n])
	reverse(s[n:])
	strconv.Atoi("11")
}


