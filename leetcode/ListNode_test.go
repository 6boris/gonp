package leetcode

import "testing"

func TestListNode_Init(t *testing.T) {
	l := ListNode{}
	data := []int{1, 2, 3, 4, 5}
	l.Init(data)
	l.Print()
}
