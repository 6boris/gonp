package leetcode

import "fmt"

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func (this *ListNode) Init(x []int) {
	head := this
	node := head
	for i := 0; i < len(x); i++ {
		node.Val = x[i]
		//	最后一个节点Next为nil
		if i != len(x)-1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
}

func (this *ListNode) Print() {

	for this != nil {
		fmt.Print(this.Val, " ")
		this = this.Next
	}
	fmt.Println()
}
