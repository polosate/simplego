package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	arr := strings.Split(text, " ")
	sum := 0
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		sum += i
	}
	fmt.Println(sum)
}


 type ListNode struct {
   Val int
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
			cur.Val = sum%10
			sum = sum/10
		} else {
			cur.Val = sum
			sum = 0
		}
		cur = cur.Next
	}

	return first
}
