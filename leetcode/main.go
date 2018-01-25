package main

import (
	"strings"
)

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var cur, first *ListNode
	sum := 0

	for l1 != nil || l2 != nil || sum != 0 {
		next := &ListNode{}
		cur = next
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum/10 > 0 {
			cur.Val = sum % 10
			sum = sum / 10
		} else {
			cur.Val = sum
			sum = 0
		}
		cur = cur.Next
	}

	return first
}

func longestPalindrome(s string) string {
	curL := [2]int{0, 0}
	maxL := [2]int{0, 0}

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > 0; j-- {
			if s[i] == s[j] {
				curL[0], curL[1] = i, j
				for k, l := i+1, j-1; k < l; k, l = k+1, l-1 {
					if s[k] != s[l] {
						curL[0], curL[1] = 0, 0
						break
					}
				}
				if curL[1]-curL[0] > maxL[1]-maxL[0] {
					maxL = curL
				}
			}
		}
	}
	if curL[1]-curL[0] > maxL[1]-maxL[0] {
		maxL = curL
	}
	return s[maxL[0] : maxL[1]+1]
}

func convert(s string, numRows int) string {
	result := make([]string, numRows)
	i := 0
	for i < len(s) {
		for j := 0; j < numRows; j++ {
			if i > len(s)-1 {
				break
			}
			result[j] +=string(s[i])
			i++
		}

		for j := numRows - 2; j > 0; j-- {
			if i > len(s)-1 {
				break
			}
			result[j] +=string(s[i])
			i++
		}
	}
	return strings.Join(result, "")
}
